# frozen_string_literal: true

# uses TrueClass/FalseClass check because rake sets `verbose`'s default to an `Object`
verbose(verbose.is_a?(TrueClass) || ENV['VERBOSE'].to_i == 1) unless verbose.is_a?(FalseClass)

require 'pathname'
require 'yaml'

require_relative 'rakelib/util'
require_relative 'rakelib/protobuf'
require_relative 'rakelib/transform_protobuf_task'
require_relative 'rakelib/generate_protobuf_task'
require_relative 'rakelib/generate_steamlang_task'

ROOT_PATH = Pathname.new(__dir__)
GO_MODULE_FILE = ROOT_PATH.join('go.mod')
RESOURCES_PATH = ROOT_PATH.join('resources')
CONFIG_FILE = RESOURCES_PATH.join('config.yml')

go_module = Util.go_module(GO_MODULE_FILE)
config = Util.load_config(CONFIG_FILE)
build_path = ROOT_PATH.join(config.fetch(:build_path))

tasklist = {
  protobuf: { transform: [], generate: [] },
  steamlang: [],
}

## Protobuf tasks

pbconfig = config.fetch(:protobuf)
proto_go_package_base = File.join(go_module, pbconfig.fetch(:go_package_base))
proto_repo_path = ROOT_PATH.join(pbconfig.fetch(:input_path))
proto_output_path = ROOT_PATH.join(pbconfig.fetch(:output_path), pbconfig.fetch(:go_package_base))
proto_patches_path = RESOURCES_PATH.join('patches', 'protobuf')
proto_transform_output_path = build_path.join('protobuf')
proto_compile_output_path = build_path.join('go')

proto_repo = ProtobufRepo.new(proto_repo_path)

pbconfig[:packages].each do |pkg_config|
  pkg_config = pkg_config.merge(
    go_package_base: proto_go_package_base,
    patches_path: proto_patches_path,
    transform_path: proto_transform_output_path,
    output_path: proto_output_path,
  )

  pkg = proto_repo.add_package(pkg_config)

  pkg.files.each do |proto_file|
    transform_file = pkg.transform_path.join(proto_file.basename)
    output_file = pkg.output_path.join(proto_file.basename.sub_ext('.pb.go'))

    transform_task = transform_protobuf(transform_file, proto_file) do |task|
      task.root_path = ROOT_PATH
    end

    generate_task = generate_protobuf(output_file, transform_task) do |task|
      task.root_path = ROOT_PATH
      task.include_path = proto_transform_output_path
      task.output_path = proto_compile_output_path
      task.go_opt = { paths: 'source_relative' }
    end

    tasklist[:protobuf][:transform] << transform_task
    tasklist[:protobuf][:generate] << generate_task
  end
end

## SteamLang tasks

slconfig = config.fetch(:steamlang)
steamlang_cmd_package = slconfig.fetch(:cmd_package)
steamlang_input_path = ROOT_PATH.join(slconfig.fetch(:input_path))
steamlang_output_path = ROOT_PATH.join(slconfig.fetch(:output_path))
steamlang_pkg_output_path = steamlang_output_path.join(slconfig.fetch(:package_name))

slconfig[:files].each_with_object(tasklist[:steamlang]) do |file, tasks|
  input_file = steamlang_input_path.join(file)
  output_file = steamlang_pkg_output_path.join(input_file.basename.sub_ext('.go'))

  tasks << generate_steamlang(output_file, input_file) do |task|
    task.root_path = ROOT_PATH
    task.cmd_package = steamlang_cmd_package
    task.go_package = slconfig.fetch(:package_name)
    task.pb_package = proto_go_package_base
  end
end

## Public tasks

task default: :generate

desc 'Clean build files'
task :clean do
  rm_rf build_path
end

namespace :generate do
  desc 'Generate protobuf files'
  task protobuf: ['generate:protobuf:transform', 'generate:protobuf:generate']

  namespace :protobuf do
    multitask transform: tasklist[:protobuf][:transform]
    multitask generate: tasklist[:protobuf][:generate]
  end

  desc 'Generate steamlang files'
  multitask steamlang: tasklist[:steamlang]
end

desc 'Generate everything'
task generate: ['generate:protobuf', 'generate:steamlang']
