require "fileutils"

#
# Release
#

class VersionedFile
  def initialize(file, regex)
    @file = file
    @regex = regex
  end

  def current_version!
    @current_version ||= matched_data![1]
  end

  def bump_version!(type)
    position = case type
               when :major
                 0
               when :minor
                 1
               when :patch
                 2
               end
    @current_version = current_version!.split(".").tap do |v|
      v[position] = v[position].to_i + 1
      # Reset consequent numbers
      ((position + 1)..2).each { |p| v[p] = 0 }
    end.join(".")
  end

  def save!
    text = File.read(@file)
    new_line = matched_data![0].gsub(matched_data![1], @current_version)
    text.gsub!(matched_data![0], new_line)

    File.open(@file, "w") { |f| f.puts text }
  end

  private

  def matched_data!
    @matched_data ||= begin
                        m = @regex.match File.read(@file)
                        raise "No version #{@regex} matched in #{@file}" unless m
                        m
                      end
  end
end

def fullpath(file)
  File.expand_path(file, File.dirname(__FILE__))
end

VERSION_FILES = {
  fullpath("README.md")           => /v(\d+.\d+.\d+)/,
  fullpath("commands/version.go") => /^const Version = "(\d+.\d+.\d+)"$/,
  fullpath(".goxc.json")          => /"PackageVersion": "(\d+.\d+.\d+)"/,
  fullpath("homebrew-gh/gh.rb")   => /VERSION = "(\d+.\d+.\d+)"/
}

class Git
  class << self
    def dirty?
      !`git status -s`.empty?
    end

    def checkout
      `git checkout .`
    end

    def commit_all(msg)
      `git commit -am "#{msg}"`
    end

    def create_tag(tag, msg)
      `git tag -a #{tag} -m "#{msg}"`
    end
  end
end

namespace :release do
  desc "Current released version"
  task :current do
    vf = VersionedFile.new(*VERSION_FILES.first)
    puts vf.current_version!
  end

  [:major, :minor, :patch].each do |type|
    desc "Release #{type} version"
    task type do
      if Git.dirty?
        puts "Please commit all changes first"
        exit 1
      end

      new_versions = VERSION_FILES.map do |file, regex|
        begin
          vf = VersionedFile.new(file, regex)
          current_version = vf.current_version!
          vf.bump_version!(type)
          vf.save!
          puts "Successfully bump #{file} from #{current_version} to #{vf.current_version!}"
          vf.current_version!
        rescue => e
          Git.checkout
          raise e
        end
      end

      require "set"
      new_versions = new_versions.to_set
      if new_versions.size != 1
        raise "More than one version found among #{VERSION_FILES}"
      end

      new_version = "v#{new_versions.first}"
      msg = "Bump version to #{new_version}"
      Git.commit_all(msg)
      Git.create_tag(new_version, msg)
    end
  end
end

#
# Build (deprecated, prefer gh_task.go)
#

module OS
  class << self
    def type
      if darwin?
        "darwin"
      elsif linux?
        "linux"
      elsif windows?
        "windows"
      else
        raise "Unknown OS type #{RUBY_PLATFORM}"
      end
    end

    def dropbox_dir
      if darwin? || linux?
        File.join ENV["HOME"], "Dropbox"
      elsif windows?
        File.join ENV["DROPBOX_DIR"]
      else
        raise "Unknown OS type #{RUBY_PLATFORM}"
      end
    end

    def windows?
      (/cygwin|mswin|mingw|bccwin|wince|emx/ =~ RUBY_PLATFORM) != nil
    end

    def darwin?
      (/darwin/ =~ RUBY_PLATFORM) != nil
    end

    def linux?
      (/linux/ =~ RUBY_PLATFORM) != nil
    end
  end
end

namespace :build do
  desc "Build for current operating system"
  task :current => [:update_goxc, :remove_build_target, :build_gh, :move_to_dropbox]

  task :update_goxc do
    puts "Updating goxc..."
    result = system "go get -u github.com/laher/goxc"
    raise "Fail to update goxc" unless result
  end

  task :remove_build_target do
    FileUtils.rm_rf fullpath("target")
  end

  task :build_gh do
    puts "Building for #{OS.type}..."
    puts `goxc -wd=. -os=#{OS.type} -c=#{OS.type}`
  end

  task :move_to_dropbox do
    vf = VersionedFile.new(*VERSION_FILES.first)
    build_dir = fullpath("target/#{vf.current_version!}-snapshot")
    dropbox_dir = File.join(OS.dropbox_dir, "Public", "gh")

    FileUtils.cp_r build_dir, dropbox_dir, :verbose => true
  end
end

#
# Tests
#

task :default => [:features]

desc "Run cucumber feature tests from Hub"
task :features do
  features = ENV.fetch("FEATURE", "features")
  sh "cucumber -f progress -t ~@wip #{features}"
end

#
# Manual
#

desc "Show man page"
task :man => "man:build" do
  exec "man man/gh.1"
end

desc "Build man pages"
task "man:build" => ["man/gh.1", "man/gh.1.html"]

extract_examples = lambda { |readme_file|
  # split readme in sections
  examples = File.read(readme_file).split(/^-{4,}$/)[6].strip
  examples.sub!(/^.+?(###)/m, '\1')  # strip intro paragraph
  examples.sub!(/\n+.+\Z/, '')       # remove last line
  examples
}

# inject examples from README file to .ronn source
source_with_examples = lambda { |source, readme|
  examples = extract_examples.call(readme)
  compiled = File.read(source)
  compiled.sub!('{{README}}', examples)
  compiled
}

# generate man page with ronn
compile_ronn = lambda { |destination, type, contents|
  File.popen("ronn --pipe --#{type} --organization=GITHUB --manual='gh Manual'", 'w+') { |io|
    io.write contents
    io.close_write
    File.open(destination, 'w') { |f| f << io.read }
  }
  abort "ronn --#{type} conversion failed" unless $?.success?
}

file "man/gh.1" => ["man/gh.1.ronn", "README.md"] do |task|
  contents = source_with_examples.call(*task.prerequisites)
  compile_ronn.call(task.name, 'roff', contents)
  compile_ronn.call("#{task.name}.html", 'html', contents)
end

file "man/gh.1.html" => ["man/gh.1.ronn", "README.md"] do |task|
  Rake::Task["man/gh.1"].invoke
end
