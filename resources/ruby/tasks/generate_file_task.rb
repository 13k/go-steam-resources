# frozen_string_literal: true

require 'pathname'
require 'rake'
require 'tempfile'

class GenerateFileTask < Rake::FileTask
  include Rake::FileUtilsExt

  def self.define_task(*args, &block)
    super.enhance { |t, _args| t.call }
  end

  attr_reader :root_path

  def root_path=(path)
    @root_path ||= Pathname.new(path)
  end

  def call
    return if executed?

    if @root_path.nil?
      raise ArgumentError, "#{self.class.name} requires `root_path` to be set"
    end

    run
  ensure
    executed!
  end

  def run
    raise NotImplementedError, "Subclasses must override this method"
  end

  protected

  def executed?
    !!@executed
  end

  def executed!
    @executed = true
  end

  def output_file
    @output_file ||= Pathname.new(name)
  end

  def output_file_from_root
    @output_file_from_root ||= output_file.relative_path_from(root_path)
  end

  def output_dir
    @output_dir ||= output_file.dirname
  end

  def input_files
    @input_files ||= prerequisite_tasks.map(&:name).map { |path| Pathname.new(path) }
  end

  def input_files_from_root
    @input_files_from_root ||= input_files.map { |p| p.relative_path_from(root_path) }
  end

  def temp_file
    @temp_files ||= []

    Tempfile.new.tap do |temp_file|
      @temp_files << temp_file
    end
  end

  def clean_temp_files
    return unless @temp_files
    @temp_files.each(&:unlink)
  end
end
