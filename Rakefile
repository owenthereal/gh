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
    @current_version = current_version!.split('.').tap do |v|
      v[position] = v[position].to_i + 1
      # Reset consequent numbers
      ((position + 1)..2).each { |p| v[p] = 0 }
    end.join('.')
  end

  def save!
    text = File.read(@file)
    new_line = matched_data![0].gsub(matched_data![1], @current_version)
    text.gsub!(matched_data![0], new_line)

    File.open(@file, 'w') { |f| f.puts text }
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

namespace :release do
  [:major, :minor, :patch].each do |type|
    desc "Release #{type} version"
    task type do
      version_file = File.expand_path('commands/version.go', File.dirname(__FILE__))
      version_regex = /^const Version = "(\d+.\d+.\d+)"$/

      readme_file = File.expand_path('README.md', File.dirname(__FILE__))
      readme_regex = /Current version is \[(\d+.\d+.\d+)\]/

      goxc_file = File.expand_path('.goxc.json', File.dirname(__FILE__))
      goxc_regex = /"PackageVersion": "(\d+.\d+.\d+)"/

      homebrew_file = File.expand_path('homebrew/gh.rb', File.dirname(__FILE__))
      homebrew_regex = /VERSION = '(\d+.\d+.\d+)'/

      {
        readme_file => readme_regex,
        goxc_file => goxc_regex,
        homebrew_file => homebrew_regex,
        version_file => version_regex
      }.each do |file, regex|
        begin
          vf = VersionedFile.new(file, regex)
          current_version = vf.current_version!
          vf.bump_version!(type)
          vf.save!
          puts "Successfully bump #{file} from #{current_version} to #{vf.current_version!}"
        rescue => e
          puts e
        end
      end
    end
  end
end
