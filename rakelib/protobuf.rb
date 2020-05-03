# frozen_string_literal: true

require 'pathname'

class ProtobufFile < Pathname
  DECL_RE = /^(?<type>enum|message|service)\s+(?<id>\w+)/.freeze

  attr_accessor :package

  def initialize(path)
    super(path)

    @package = nil
  end

  def symbols
    @symbols ||= each_line.grep(DECL_RE) do
      {
        type: Regexp.last_match[:type],
        id: Regexp.last_match[:id],
      }
    end
  end
end

class ProtobufPackage
  attr_reader :repo

  def initialize(repo, config)
    @repo = repo
    @config = config
  end

  def name
    @name ||= @config.fetch(:name)
  end

  alias path name

  def pb_package
    @pb_package ||= name.gsub('/', '.')
  end

  def go_package
    @go_package ||= File.join(@config.fetch(:go_package_base), name)
  end

  def patches_path
    @patches_path ||= @config.fetch(:transform_path).join(name)
  end

  def transform_path
    @transform_path ||= @config.fetch(:transform_path).join(name)
  end

  def output_path
    @output_path ||= @config.fetch(:output_path).join(name)
  end

  def files
    @files ||= @config.fetch(:files).flat_map { |patt| @repo.glob(patt, exclude: excludes) }
  end

  protected

  def excludes
    @excludes ||= @config.fetch(:exclude, [])
  end
end

class ProtobufRepo
  attr_reader :path

  def initialize(path)
    @path = path
    @files = {}
    @packages = {}
  end

  def glob(pattern, exclude: [])
    @path
      .glob(pattern)
      .reject { |p| exclude.any? { |patt| p.fnmatch?(@path.join(patt).to_s, File::FNM_PATHNAME) } }
      .map { |p| ProtobufFile.new(p) }
  end

  def add_package(config)
    ProtobufPackage.new(self, config).tap do |pkg|
      @packages[pkg.name] = pkg

      pkg.files.each do |f|
        f.package = pkg
        @files[f.to_s] = f
      end
    end
  end

  def [](path)
    @files[path.to_s]
  end
end
