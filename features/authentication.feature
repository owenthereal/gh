Feature: OAuth authentication
  Background:
    Given I am in "dotfiles" git repo

  Scenario: Ask for username & password, create authorization
    Given the GitHub API server:
      """
      require 'rack/auth/basic'
      get('/authorizations') { json [] }
      post('/authorizations') {
        auth = Rack::Auth::Basic::Request.new(env)
        unless auth.credentials == %w[mislav kitty]
          halt 401, json(:error => 'error')
        end
        assert :scopes => ['repo']
        json :token => 'OTOKEN'
      }
      get('/user') {
        unless request.env['HTTP_AUTHORIZATION'] == 'token OTOKEN'
          halt 401, json(:error => 'error')
        end
        json :login => 'MiSlAv'
      }
      post('/user/repos') {
        unless request.env['HTTP_AUTHORIZATION'] == 'token OTOKEN'
          halt 401, json(:error => 'error')
        end
        json :full_name => 'mislav/dotfiles'
      }
      """
    When I run `hub create` interactively
    When I type "mislav"
    And I type "kitty"
    Then the output should contain "github.com username:"
    And the output should contain "github.com password for mislav (never stored):"
    And the exit status should be 0
    And the file "../home/.config/gh" should contain "mislav"
    And the file "../home/.config/gh" should contain "OTOKEN"
    #And the file "../home/.config/gh" should have mode "0600"

  Scenario: Ask for username & password, re-use existing authorization
    Given the GitHub API server:
      """
      require 'rack/auth/basic'
      get('/authorizations') {
        auth = Rack::Auth::Basic::Request.new(env)
        unless auth.credentials == %w[mislav kitty]
          halt 401, json(:error => 'error')
        end

        json [
          {:token => 'SKIPPD', :app => {:url => 'http://example.com'}},
          {:token => 'OTOKEN', :app => {:url => 'http://owenou.com/gh'}}
        ]
      }
      get('/user') {
        json :login => 'mislav'
      }
      post('/user/repos') {
        json :full_name => 'mislav/dotfiles'
      }
      """
    When I run `hub create` interactively
    When I type "mislav"
    And I type "kitty"
    Then the output should contain "github.com password for mislav (never stored):"
    And the exit status should be 0
    And the file "../home/.config/gh" should contain "OTOKEN"

  Scenario: Credentials from GITHUB_USER & GITHUB_PASSWORD
    Given the GitHub API server:
      """
      require 'rack/auth/basic'
      get('/authorizations') {
        auth = Rack::Auth::Basic::Request.new(env)
        unless auth.credentials == %w[mislav kitty]
          halt 401, json(:error => 'error')
        end
        json [
          {:token => 'OTOKEN', :app => {:url => 'http://owenou.com/gh'}}
        ]
      }
      get('/user') {
        json :login => 'mislav'
      }
      post('/user/repos') {
        json :full_name => 'mislav/dotfiles'
      }
      """
    Given $GITHUB_USER is "mislav"
    And $GITHUB_PASSWORD is "kitty"
    When I successfully run `hub create`
    Then the output should not contain "github.com password for mislav"
    And the file "../home/.config/gh" should contain "OTOKEN"

  Scenario: Wrong password
    Given the GitHub API server:
      """
      require 'rack/auth/basic'
      get('/authorizations') {
        auth = Rack::Auth::Basic::Request.new(env)
        unless auth.credentials == %w[mislav kitty]
          halt 401, json(:error => 'auth error')
        end
      }
      """
    When I run `hub create` interactively
    When I type "mislav"
    And I type "WRONG"
    Then the stderr should contain "401 - Error: auth error"
    And the exit status should be 1
    #And the file "../home/.config/gh" should not exist

  Scenario: Two-factor authentication, create authorization
    Given the GitHub API server:
      """
      require 'rack/auth/basic'
      get('/authorizations') {
        auth = Rack::Auth::Basic::Request.new(env)
        unless auth.credentials == %w[mislav kitty]
          halt 401, json(:error => 'error')
        end

        if request.env['HTTP_X_GITHUB_OTP'] != "112233"
          response.headers['X-GitHub-OTP'] = "required; application"
          halt 401, json(:error => 'two-factor authorization OTP code')
        end

        json [
        ]
      }
      post('/authorizations') {
        auth = Rack::Auth::Basic::Request.new(env)
        unless auth.credentials == %w[mislav kitty]
          halt 401, json(:error => 'error')
        end

        unless params[:scopes]
          halt 412, json(:error => 'error')
        end

        if request.env['HTTP_X_GITHUB_OTP'] != "112233"
          response.headers['X-GitHub-OTP'] = "required; application"
          halt 401, json(:error => 'two-factor authentication OTP code')
        end

        json :token => 'OTOKEN'
      }

      get('/user') {
        json :login => 'mislav'
      }

      post('/user/repos') {
        json :full_name => 'mislav/dotfiles'
      }
      """
    When I run `hub create` interactively
    When I type "mislav"
    And I type "kitty"
    And I type "112233"
    Then the output should contain "github.com password for mislav (never stored):"
    Then the output should contain "two-factor authentication code:"
    And the exit status should be 0
    And the file "../home/.config/gh" should contain "OTOKEN"

  Scenario: Two-factor authentication, re-use existing authorization
    Given the GitHub API server:
      """
      token = 'OTOKEN'
      post('/authorizations') {
        assert_basic_auth 'mislav', 'kitty'
        token << 'SMS'
        halt 412, json(:error => 'error')
      }
      get('/authorizations') {
        assert_basic_auth 'mislav', 'kitty'
        if request.env['HTTP_X_GITHUB_OTP'] != "112233"
          response.headers['X-GitHub-OTP'] = "required; application"
          halt 401, json(:error => 'error')
        end
        json [ {
          :token => token,
          :app => {:url => 'http://owenou.com/gh'}
          } ]
      }
      get('/user') {
        json :login => 'mislav'
      }
      post('/user/repos') {
        json :full_name => 'mislav/dotfiles'
      }
      """
    When I run `hub create` interactively
    When I type "mislav"
    And I type "kitty"
    And I type "112233"
    Then the output should contain "github.com password for mislav (never stored):"
    Then the output should contain "two-factor authentication code:"
    And the exit status should be 0
    And the file "../home/.config/gh" should contain "OTOKEN"

  @wip
  Scenario: Special characters in username & password
    Given the GitHub API server:
      """
      get('/authorizations') { json [] }
      post('/authorizations') {
        assert_basic_auth 'mislav@example.com', 'my pass@phrase ok?'
        json :token => 'OTOKEN'
      }
      get('/user') {
        json :login => 'mislav'
      }
      post('/user/repos') {
        unless request.env['HTTP_AUTHORIZATION'] == 'token OTOKEN'
          halt 401, json(:error => 'error')
        end
        json :full_name => 'mislav/dotfiles'
      }
      """
    When I run `hub create` interactively
    When I type "mislav@example.com"
    And I type "my pass@phrase ok?"
    Then the output should contain "github.com password for mislav@example.com (never stored):"
    And the exit status should be 0
    And the file "../home/.config/gh" should contain "mislav"
    And the file "../home/.config/gh" should contain "OTOKEN"
