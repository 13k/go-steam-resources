# frozen_string_literal: true

# uses TrueClass/FalseClass check because rake sets `verbose`'s default to an `Object`
verbose(verbose.is_a?(TrueClass) || ENV['VERBOSE'].to_i == 1) unless verbose.is_a?(FalseClass)

require 'pathname'
require 'yaml'

ROOT_PATH = Pathname.new(__dir__)
RESOURCES_PATH = ROOT_PATH.join('resources')

require_relative RESOURCES_PATH.join('ruby/logging')
require_relative RESOURCES_PATH.join('ruby/util')
require_relative RESOURCES_PATH.join('ruby/tasks/transform_protobuf_task')
require_relative RESOURCES_PATH.join('ruby/tasks/generate_protobuf_task')
require_relative RESOURCES_PATH.join('ruby/tasks/generate_steamlang_task')

GO_MODULE_FILE = ROOT_PATH.join('go.mod')
CONFIG_FILE = RESOURCES_PATH.join('config.yml')

go_module = parse_go_module(GO_MODULE_FILE)
config = load_config(CONFIG_FILE)
build_path = ROOT_PATH.join(config.fetch(:build_path))

tasklist = {
  protobuf: { transform: [], generate: [] },
  steamlang: [],
}

## Protobuf tasks

pbconfig = config.fetch(:protobuf)
protobuf_pkg_base_import_path = File.join(go_module, pbconfig.fetch(:base_package_name))
protobuf_input_path = ROOT_PATH.join(pbconfig[:input_path])
protobuf_output_path = ROOT_PATH.join(pbconfig[:output_path])
protobuf_pkg_output_path = protobuf_output_path.join(pbconfig.fetch(:base_package_name))
protobuf_patched_output_path = build_path.join('protobufs')
protobuf_generated_output_path = build_path.join('go')
protobuf_patches_path = ROOT_PATH.join(pbconfig.fetch(:patches_path))

pbconfig[:packages].each_with_object(tasklist[:protobuf]) do |package, tasks|
  pkg_import_path = File.join(protobuf_pkg_base_import_path, package[:name])
  pkg_patches_path = protobuf_patches_path.join(package[:name])
  pkg_patched_path = protobuf_patched_output_path.join(package[:name])
  pkg_output_path = protobuf_pkg_output_path.join(package[:name])

  excludes = package
    .fetch(:exclude_files, [])
    .map { |pattern| protobuf_input_path.join(pattern).to_s }

  package[:files].flat_map do |pattern|
    protobuf_input_path
      .glob(pattern)
      .reject { |file| excludes.any? { |excl_pattern| file.fnmatch?(excl_pattern) } }
      .map do |input_file|
        transformed_file = pkg_patched_path.join(input_file.basename)
        output_file = pkg_output_path.join(input_file.basename.sub_ext('.pb.go'))

        transform = TransformProtobufTask.define_task(transformed_file => input_file) do |task|
          task.root_path = ROOT_PATH
          task.patches_path = pkg_patches_path
          task.go_package = pkg_import_path
        end

        generate = GenerateProtobufTask.define_task(output_file => transform) do |task|
          task.root_path = ROOT_PATH
          task.import_path = pkg_import_path
          task.output_path = protobuf_generated_output_path
        end

        tasks[:transform] << transform
        tasks[:generate] << generate
      end
  end
end

## SteamLang tasks

slconfig = config.fetch(:steamlang)
steamlang_gen_go_file = ROOT_PATH.join(slconfig.fetch(:gen_go_file))
steamlang_input_path = ROOT_PATH.join(slconfig.fetch(:input_path))
steamlang_output_path = ROOT_PATH.join(slconfig.fetch(:output_path))
steamlang_pkg_output_path = steamlang_output_path.join(slconfig.fetch(:package_name))

slconfig[:files].each_with_object(tasklist[:steamlang]) do |file, tasks|
  input_file = steamlang_input_path.join(file)
  output_file = steamlang_pkg_output_path.join(input_file.basename.sub_ext('.go'))

  tasks << GenerateSteamLangTask.define_task(output_file => input_file) do |task|
    task.root_path = ROOT_PATH
    task.gen_go_file = steamlang_gen_go_file
    task.go_package = slconfig.fetch(:package_name)
    task.protobuf_package = protobuf_pkg_base_import_path
  end
end

## Public tasks

task default: :generate

desc "Clean build files"
task :clean do
  rm_rf build_path
end

namespace :generate do
  desc "Generate protobuf files"
  task protobuf: ['generate:protobuf:transform', 'generate:protobuf:generate']

  namespace :protobuf do
    multitask transform: tasklist[:protobuf][:transform]
    multitask generate: tasklist[:protobuf][:generate]
  end

  desc "Generate steamlang files"
  multitask steamlang: tasklist[:steamlang]
end

desc "Generate everything"
task generate: ['generate:protobuf', 'generate:steamlang']
