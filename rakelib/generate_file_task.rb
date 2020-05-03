# frozen_string_literal: true

require 'pathname'
require 'rake'
require 'tempfile'

require_relative 'util'

class GenerateFileTask < Rake::FileTask
  include Rake::FileUtilsExt

  def self.define_task(output_file, *inputs, &block)
    task = super(output_file => inputs, &block).tap do |t|
      t.output_file = output_file
      t.inputs = inputs
    end

    task.enhance { |t, _| t.call }
  end

  attr_accessor :root_path
  attr_reader :output_file, :inputs

  def output_file=(output_file)
    @output_file = require_pathname!(output_file)
  end

  def inputs=(inputs)
    @inputs = Array(inputs).map { |i| validate_input!(i) }
  end

  def needed?
    !@output_file.exist? || out_of_date?(timestamp) || @application.options.build_all
  end

  def timestamp
    @output_file.exist? ? @output_file.mtime : Rake::LATE
  end

  def call
    if @root_path.nil?
      raise ArgumentError, format(
        '%<task_class>s requires `root_path` to be set',
        task_class: self.class.name,
      )
    end

    unless @root_path.is_a?(Pathname)
      raise ArgumentError, format(
        '%<task_class>s requires `root_path` to be a Pathname object, got %<root_path_class>s',
        task_class: self.class.name,
        root_path_class: @root_path.class.name,
      )
    end

    run
  end

  def run
    raise NotImplementedError, 'Subclasses must override this method'
  end

  protected

  def output_file_from_root
    @output_file_from_root ||= @output_file.relative_path_from(@root_path)
  end

  def output_dir
    @output_dir ||= @output_file.dirname
  end

  def input_files
    @input_files ||= @inputs.map do |input|
      case input
      when GenerateFileTask
        input.output_file
      when Rake::Task
        Pathname.new(input.name)
      else
        input
      end
    end
  end

  def input_files_from_root
    @input_files_from_root ||= input_files.map { |p| p.relative_path_from(@root_path) }
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

  def require_command!(cmd)
    Util.which!(cmd)
  end

  def require_pathname!(arg)
    unless arg.is_a?(Pathname)
      raise ArgumentError, format(
        '%<task>s requires output and inputs to be Pathname objects, got %<actual>s',
        task: self.class.name,
        actual: arg.class.name,
      )
    end

    arg
  end

  def validate_input!(arg)
    return arg if arg.is_a?(Rake::FileTask)

    require_pathname!(arg)
  end
end
