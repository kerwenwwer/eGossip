// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64

package bpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfMetadata struct {
	Metadata   [256]int8
	UpdateTime int64
}

type bpfTargets struct {
	TargetList [64]struct {
		Ip   uint32
		Port uint16
		Mac  [6]int8
	}
	MaxCount uint16
	_        [2]byte
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	Fastbroadcast *ebpf.ProgramSpec `ebpf:"fastbroadcast"`
	XdpSockProg   *ebpf.ProgramSpec `ebpf:"xdp_sock_prog"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	MetadataMap *ebpf.MapSpec `ebpf:"metadata_map"`
	NodelistMap *ebpf.MapSpec `ebpf:"nodelist_map"`
	QidconfMap  *ebpf.MapSpec `ebpf:"qidconf_map"`
	TargetsMap  *ebpf.MapSpec `ebpf:"targets_map"`
	XsksMap     *ebpf.MapSpec `ebpf:"xsks_map"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	MetadataMap *ebpf.Map `ebpf:"metadata_map"`
	NodelistMap *ebpf.Map `ebpf:"nodelist_map"`
	QidconfMap  *ebpf.Map `ebpf:"qidconf_map"`
	TargetsMap  *ebpf.Map `ebpf:"targets_map"`
	XsksMap     *ebpf.Map `ebpf:"xsks_map"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.MetadataMap,
		m.NodelistMap,
		m.QidconfMap,
		m.TargetsMap,
		m.XsksMap,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	Fastbroadcast *ebpf.Program `ebpf:"fastbroadcast"`
	XdpSockProg   *ebpf.Program `ebpf:"xdp_sock_prog"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.Fastbroadcast,
		p.XdpSockProg,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel.o
var _BpfBytes []byte
