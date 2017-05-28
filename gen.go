// +build ignore

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	packageNs                  = "github.com/13k/go-steam-resources"
	baseProtobufsInputRelPath  = "SteamKit/Resources/Protobufs"
	baseProtobufsVendorRelPath = "protobuf/vendor"
)

var (
	basePath                string
	baseProtobufsInputPath  string
	baseProtobufsOutputPath string
	baseProtobufsVendorPath string

	pbPackagesByName = map[string]*pbPackage{}

	pbPackages = []*pbPackage{
		&pbPackage{"google/protobuf", "google/protobuf", true, nil, []string{
			"descriptor.proto",
		}},
		&pbPackage{"protobuf", "steamclient", false, nil, []string{
			"steammessages_base.proto",
			"content_manifest.proto",
			"encrypted_app_ticket.proto",
			"steammessages_clientserver.proto",
			"steammessages_clientserver_2.proto",
		}},
		&pbPackage{"protobuf/unified", "steamclient", false, nil, []string{
			"steammessages_unified_base.steamclient.proto",
			"steammessages_broadcast.steamclient.proto",
			"steammessages_cloud.steamclient.proto",
			"steammessages_credentials.steamclient.proto",
			"steammessages_depotbuilder.steamclient.proto",
			"steammessages_deviceauth.steamclient.proto",
			"steammessages_econ.steamclient.proto",
			"steammessages_gamenotifications.steamclient.proto",
			"steammessages_gameservers.steamclient.proto",
			"steammessages_inventory.steamclient.proto",
			"steammessages_linkfilter.steamclient.proto",
			"steammessages_offline.steamclient.proto",
			"steammessages_parental.steamclient.proto",
			"steammessages_partnerapps.steamclient.proto",
			"steammessages_player.steamclient.proto",
			"steammessages_publishedfile.steamclient.proto",
			"steammessages_secrets.steamclient.proto",
			"steammessages_twofactor.steamclient.proto",
			"steammessages_video.steamclient.proto",
		}},
		&pbPackage{"protobuf/unified/steamworks", "dota", false, nil, []string{
			"steammessages_oauth.steamworkssdk.proto",
		}},
		&pbPackage{"protobuf/gc", "dota", false, nil, []string{
			"steammessages.proto",
			"gcsystemmsgs.proto",
			"base_gcmessages.proto",
			"gcsdk_gcmessages.proto",
			"econ_gcmessages.proto",
		}},
		&pbPackage{"protobuf/gc/csgo", "csgo", false, nil, []string{
			"cstrike15_gcmessages.proto",
		}},
		&pbPackage{"protobuf/gc/dota", "dota", false, []string{"protobuf/gc"}, []string{
			"network_connection.proto",
			"dota_gcmessages_common.proto",
			"dota_gcmessages_client.proto",
			"dota_gcmessages_client_fantasy.proto",
			"dota_gcmessages_server.proto",
		}},
		&pbPackage{"protobuf/gc/tf2", "tf", false, nil, []string{
			"base_gcmessages.proto",
			"tf_gcmessages.proto",
		}},
	}
)

func init() {
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	} else {
		if p, err := os.Getwd(); err != nil {
			panic(err)
		} else {
			basePath = p
		}
	}

	baseProtobufsInputPath = filepath.Join(basePath, baseProtobufsInputRelPath)
	baseProtobufsOutputPath = basePath
	baseProtobufsVendorPath = filepath.Join(baseProtobufsOutputPath, baseProtobufsVendorRelPath)

	for _, pkg := range pbPackages {
		pbPackagesByName[pkg.name] = pkg
	}
}

type pbPackage struct {
	name         string
	dir          string
	vendor       bool
	dependencies []string
	files        []string
}

func (p *pbPackage) ImportPath() string {
	var base string

	if p.vendor {
		base = ""
	} else {
		base = packageNs
	}

	return path.Join(base, p.name)
}

func (p *pbPackage) Dependencies() []*pbPackage {
	var dependencies []*pbPackage

	for _, dep := range p.dependencies {
		if pkg, ok := pbPackagesByName[dep]; ok {
			dependencies = append(dependencies, pkg)
		}
	}

	return dependencies
}

func (p *pbPackage) OutputDir() string {
	var basePath string

	if p.vendor {
		basePath = baseProtobufsVendorPath
	} else {
		basePath = baseProtobufsOutputPath
	}

	return filepath.Join(basePath, p.name)
}

func (p *pbPackage) ProtobufsDir() string {
	return filepath.Join(baseProtobufsInputPath, p.dir)
}

func (p *pbPackage) IncludeDir() string {
	return p.dir
}

func (p *pbPackage) Build() error {
	if err := os.MkdirAll(p.OutputDir(), os.ModePerm); err != nil {
		return err
	}

	if err := p.compile(); err != nil {
		return err
	}

	return nil
}

func (p *pbPackage) compile() error {
	opts := p.compileOptions()
	goOut := fmt.Sprintf("--go_out=%s:%s", strings.Join(opts, ","), p.OutputDir())

	cmd := []string{
		"protoc",
		"-I", ".",
		"-I", p.IncludeDir(),
		goOut,
	}

	cmd = append(cmd, p.files...)

	err, stdout, stderr := run(cmd, p.ProtobufsDir())

	if err != nil {
		fmt.Fprintf(os.Stdout, stdout)
		fmt.Fprintf(os.Stderr, stderr)
	}

	return err
}

func (p *pbPackage) compileOptions() []string {
	var opts []string

	pkgs := []*pbPackage{p}
	pkgs = append(pkgs, p.Dependencies()...)

	for _, pkg := range pkgs {
		for _, file := range pkg.files {
			option := fmt.Sprintf("M%s=%s", file, pkg.ImportPath())
			opts = append(opts, option)
		}
	}

	opts = append(opts, fmt.Sprintf("import_path=%s", p.ImportPath()))

	return opts
}

func run(args []string, wd string) (error, string, string) {
	cmd := exec.Command(args[0], args[1:]...)
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = wd
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func genProto() {
	for _, pkg := range pbPackages {
		if err := pkg.Build(); err != nil {
			panic(err)
		}
	}
}

func main() {
	genProto()
}
