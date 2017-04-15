Feature: hub fork
  Background:
    Given I am in "dotfiles" git repo
    And the "origin" remote has url "git://github.com/evilchelu/dotfiles.git"
    And I am "mislav" on github.com with OAuth token "OTOKEN"

  Scenario: Fork the repository
    Given the GitHub API server:
      """
      before do
        unless request.env['HTTP_AUTHORIZATION'] == 'token OTOKEN'
          status 401
          json :message => "I haz fail!"
        end
      end

      get('/repos/mislav/dotfiles') do
        status 404
        json :message => "I haz fail!"
      end

      post('/repos/evilchelu/dotfiles/forks') do
        json :html_url => "https://github.com/mislav/coral/pull/12"
      end
      """
    When I successfully run `hub fork`
    Then the output should contain exactly "new remote: mislav\n"
    And "git remote add -f mislav git://github.com/evilchelu/dotfiles.git" should be run
    And "git remote set-url mislav git@github.com:mislav/dotfiles.git" should be run
    And the url for "mislav" should be "git@github.com:mislav/dotfiles.git"

  Scenario: Fork the repository when origin URL is private
    Given the "origin" remote has url "git@github.com:evilchelu/dotfiles.git"
    Given the GitHub API server:
      """
      before do
        unless request.env['HTTP_AUTHORIZATION'] == 'token OTOKEN'
          status 401
          json :message => "I haz fail!"
        end
      end

      get('/repos/mislav/dotfiles') do
        status 404
        json :message => "I haz fail!"
      end

      post('/repos/evilchelu/dotfiles/forks') do
        json :html_url => "https://github.com/mislav/coral/pull/12"
      end
      """
    When I successfully run `hub fork`
    Then the output should contain exactly "new remote: mislav\n"
    And "git remote add -f mislav ssh://git@github.com/evilchelu/dotfiles.git" should be run
    And "git remote set-url mislav git@github.com:mislav/dotfiles.git" should be run
    And the url for "mislav" should be "git@github.com:mislav/dotfiles.git"

  Scenario: --no-remote
    Given the GitHub API server:
      """
      post('/repos/evilchelu/dotfiles/forks') do
        json :repo => "repo"
      end
      """
    When I successfully run `hub fork --no-remote`
    Then there should be no output
    And there should be no "mislav" remote

  Scenario: Fork failed
    Given the GitHub API server:
      """
      post('/repos/evilchelu/dotfiles/forks') do
        status 500
        json(:error => "I haz fail!")
      end
      """
    When I run `hub fork`
    Then the exit status should be 1
    And the stderr should contain exactly:
      """
      500 - Error: I haz fail!\n
      """
    And there should be no "mislav" remote

  Scenario: Unrelated fork already exists
    Given the GitHub API server:
      """
      get('/repos/mislav/dotfiles') {
        halt 406 unless request.env['HTTP_ACCEPT'] == 'application/vnd.github.v3+json'
        json :parent => { :html_url => 'https://github.com/unrelated/dotfiles' }
      }
      """
    When I run `hub fork`
    Then the exit status should be 1
    And the stderr should contain exactly:
      """
      Error creating fork: mislav/dotfiles already exists on github.com\n
      """
    And there should be no "mislav" remote

Scenario: Related fork already exists
    Given the GitHub API server:
      """
      get('/repos/mislav/dotfiles') {
        json :parent => { :html_url => 'https://github.com/evilchelu/dotfiles' }
      }
      """
    When I run `hub fork`
    Then the exit status should be 0
    And the url for "mislav" should be "git@github.com:mislav/dotfiles.git"

  Scenario: Invalid OAuth token
    Given the GitHub API server:
      """
      before do
        unless request.env['HTTP_AUTHORIZATION'] == 'token OTOKEN'
          halt 401, json(:message => "I haz fail!")
        end
      end
      """
    And I am "mislav" on github.com with OAuth token "WRONGTOKEN"
    When I run `hub fork`
    Then the exit status should be 1
    And the stderr should contain exactly:
      """
      401 - I haz fail!
      """

  Scenario: HTTPS is preferred
    Given the GitHub API server:
      """
      post('/repos/evilchelu/dotfiles/forks') { json :repo => 'repo' }
      """
    And HTTPS is preferred
    When I successfully run `hub fork`
    Then the output should contain exactly "new remote: mislav\n"
    And the url for "mislav" should be "https://github.com/mislav/dotfiles.git"

  Scenario: Not in repo
    Given the current dir is not a repo
    When I run `hub fork`
    Then the exit status should be 1
    And the stderr should contain "Aborted: the origin remote doesn't point to a GitHub repository"

  Scenario: Unknown host
    Given the "origin" remote has url "git@git.my.org:evilchelu/dotfiles.git"
    When I run `hub fork`
    Then the exit status should be 1
    And the stderr should contain:
      """
      Aborted: the origin remote doesn't point to a GitHub repository
      """

  Scenario: Enterprise fork
    Given the GitHub API server:
      """
      before do
        unless request.env['HTTP_AUTHORIZATION'] == 'token FITOKEN'
          status 401 
          json :error => 'error'
        end
      end
      post('/api/v3/repos/evilchelu/dotfiles/forks') { json :repo => 'repo' }
      """
    And the "origin" remote has url "git@git.my.org:evilchelu/dotfiles.git"
    And I am "mislav" on git.my.org with OAuth token "FITOKEN"
    And "git.my.org" is a whitelisted Enterprise host
    When I successfully run `hub fork`
    Then the url for "mislav" should be "git@git.my.org:mislav/dotfiles.git"
