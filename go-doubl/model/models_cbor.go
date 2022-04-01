// Code generated by github.com/daotl/cbor-gen. DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"math"
	"sort"

	cbg "github.com/daotl/cbor-gen"
	cid "github.com/ipfs/go-cid"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *Transaction) InitNilEmbeddedStruct() {
	if t != nil {
	}
}

var lengthBufTransaction = []byte{135}

func (t *Transaction) MarshalCBOR(w io.Writer) (n int, err error) {
	if t == nil {
		return w.Write(cbg.CborNull)
	}
	t.InitNilEmbeddedStruct()
	if n_, err := w.Write(lengthBufTransaction); err != nil {
		return n_, err
	} else {
		n += n_
	}

	scratch := make([]byte, 9)

	// t.Type (model.TransactionType) (uint8)
	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Type)); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.From (bytes.HexBytes) (slice)
	if len(t.From) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.From was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.From))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.From[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.Nonce (uint64) (uint64)

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Nonce)); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.To (bytes.HexBytes) (slice)
	if len(t.To) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.To was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.To))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.To[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.Data ([]uint8) (slice)
	if len(t.Data) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.Data was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Data))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.Data[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.Extra ([]uint8) (slice)
	if len(t.Extra) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.Extra was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Extra))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.Extra[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.Sig (crpt.Signature) (slice)
	if len(t.Sig) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.Sig was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Sig))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.Sig[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}
	return n, nil
}

func (t *Transaction) UnmarshalCBOR(r io.Reader) (int, error) {
	bytesRead := 0
	*t = Transaction{}
	t.InitNilEmbeddedStruct()

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, read, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read
	if maj != cbg.MajArray {
		return bytesRead, fmt.Errorf("cbor input should be of type array")
	}

	if extra != 7 {
		return bytesRead, fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Type (model.TransactionType) (uint8)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read
	if maj != cbg.MajUnsignedInt {
		return bytesRead, fmt.Errorf("wrong type for uint8 field")
	}
	if extra > math.MaxUint8 {
		return bytesRead, fmt.Errorf("integer in input was too large for uint8 field")
	}
	t.Type = TransactionType(extra)
	// t.From (bytes.HexBytes) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.From: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.From = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.From[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	// t.Nonce (uint64) (uint64)

	{

		maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return bytesRead, err
		}
		bytesRead += read
		if maj != cbg.MajUnsignedInt {
			return bytesRead, fmt.Errorf("wrong type for uint64 field")
		}
		t.Nonce = uint64(extra)

	}
	// t.To (bytes.HexBytes) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.To: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.To = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.To[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	// t.Data ([]uint8) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.Data: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Data = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.Data[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	// t.Extra ([]uint8) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.Extra: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Extra = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.Extra[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	// t.Sig (crpt.Signature) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.Sig: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Sig = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.Sig[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	return bytesRead, nil
}

func (t *BlockHeader) InitNilEmbeddedStruct() {
	if t != nil {
	}
}

var lengthBufBlockHeader = []byte{137}

func (t *BlockHeader) MarshalCBOR(w io.Writer) (n int, err error) {
	if t == nil {
		return w.Write(cbg.CborNull)
	}
	t.InitNilEmbeddedStruct()
	if n_, err := w.Write(lengthBufBlockHeader); err != nil {
		return n_, err
	} else {
		n += n_
	}

	scratch := make([]byte, 9)

	// t.Creator (bytes.HexBytes) (slice)
	if len(t.Creator) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.Creator was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Creator))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.Creator[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.Time (model.Timestamp) (uint64)

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Time)); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.PrevHashes ([][]uint8) (slice)
	if len(t.PrevHashes) > cbg.MaxLength {
		return n, xerrors.Errorf("Slice value in field t.PrevHashes was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.PrevHashes))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}
	for _, v := range t.PrevHashes {
		if len(v) > cbg.ByteArrayMaxLen {
			return n, xerrors.Errorf("Byte array in field v was too long")
		}

		if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(v))); err != nil {
			return n + n_, err
		} else {
			n += n_
		}

		if n_, err := w.Write(v[:]); err != nil {
			return n + n_, err
		} else {
			n += n_
		}
	}

	// t.Height (model.BlockHeight) (uint64)

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Height)); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.TxRoot ([]uint8) (slice)
	if len(t.TxRoot) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.TxRoot was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.TxRoot))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.TxRoot[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.TxCount (uint64) (uint64)

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.TxCount)); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.AppHash ([]uint8) (slice)
	if len(t.AppHash) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.AppHash was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.AppHash))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.AppHash[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.Extra ([]uint8) (slice)
	if len(t.Extra) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.Extra was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Extra))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.Extra[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.Sig (crpt.Signature) (slice)
	if len(t.Sig) > cbg.ByteArrayMaxLen {
		return n, xerrors.Errorf("Byte array in field t.Sig was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Sig))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	if n_, err := w.Write(t.Sig[:]); err != nil {
		return n + n_, err
	} else {
		n += n_
	}
	return n, nil
}

func (t *BlockHeader) UnmarshalCBOR(r io.Reader) (int, error) {
	bytesRead := 0
	*t = BlockHeader{}
	t.InitNilEmbeddedStruct()

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, read, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read
	if maj != cbg.MajArray {
		return bytesRead, fmt.Errorf("cbor input should be of type array")
	}

	if extra != 9 {
		return bytesRead, fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Creator (bytes.HexBytes) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.Creator: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Creator = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.Creator[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	// t.Time (model.Timestamp) (uint64)

	{

		maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return bytesRead, err
		}
		bytesRead += read
		if maj != cbg.MajUnsignedInt {
			return bytesRead, fmt.Errorf("wrong type for uint64 field")
		}
		t.Time = Timestamp(extra)

	}
	// t.PrevHashes ([][]uint8) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.MaxLength {
		return bytesRead, fmt.Errorf("t.PrevHashes: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return bytesRead, fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.PrevHashes = make([][]uint8, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error

			maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return bytesRead, err
			}
			bytesRead += read

			if extra > cbg.ByteArrayMaxLen {
				return bytesRead, fmt.Errorf("t.PrevHashes[i]: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return bytesRead, fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.PrevHashes[i] = make([]uint8, extra)
			}

			if read, err := io.ReadFull(br, t.PrevHashes[i][:]); err != nil {
				return bytesRead, err
			} else {
				bytesRead += read
			}
		}
	}

	// t.Height (model.BlockHeight) (uint64)

	{

		maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return bytesRead, err
		}
		bytesRead += read
		if maj != cbg.MajUnsignedInt {
			return bytesRead, fmt.Errorf("wrong type for uint64 field")
		}
		t.Height = BlockHeight(extra)

	}
	// t.TxRoot ([]uint8) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.TxRoot: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.TxRoot = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.TxRoot[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	// t.TxCount (uint64) (uint64)

	{

		maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return bytesRead, err
		}
		bytesRead += read
		if maj != cbg.MajUnsignedInt {
			return bytesRead, fmt.Errorf("wrong type for uint64 field")
		}
		t.TxCount = uint64(extra)

	}
	// t.AppHash ([]uint8) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.AppHash: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.AppHash = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.AppHash[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	// t.Extra ([]uint8) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.Extra: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Extra = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.Extra[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	// t.Sig (crpt.Signature) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.ByteArrayMaxLen {
		return bytesRead, fmt.Errorf("t.Sig: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return bytesRead, fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Sig = make([]uint8, extra)
	}

	if read, err := io.ReadFull(br, t.Sig[:]); err != nil {
		return bytesRead, err
	} else {
		bytesRead += read
	}
	return bytesRead, nil
}

func (t *Block) InitNilEmbeddedStruct() {
	if t != nil {
	}
}

var lengthBufBlock = []byte{130}

func (t *Block) MarshalCBOR(w io.Writer) (n int, err error) {
	if t == nil {
		return w.Write(cbg.CborNull)
	}
	t.InitNilEmbeddedStruct()
	if n_, err := w.Write(lengthBufBlock); err != nil {
		return n_, err
	} else {
		n += n_
	}

	scratch := make([]byte, 9)

	// t.Header (model.BlockHeader) (struct)
	if n_, err := t.Header.MarshalCBOR(w); err != nil {
		return n + n_, err
	} else {
		n += n_
	}

	// t.Txs (model.TransactionSlice) (slice)
	if len(t.Txs) > cbg.MaxLength {
		return n, xerrors.Errorf("Slice value in field t.Txs was too long")
	}

	if n_, err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Txs))); err != nil {
		return n + n_, err
	} else {
		n += n_
	}
	for _, v := range t.Txs {
		if n_, err := v.MarshalCBOR(w); err != nil {
			return n + n_, err
		} else {
			n += n_
		}
	}
	return n, nil
}

func (t *Block) UnmarshalCBOR(r io.Reader) (int, error) {
	bytesRead := 0
	*t = Block{}
	t.InitNilEmbeddedStruct()

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, read, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read
	if maj != cbg.MajArray {
		return bytesRead, fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return bytesRead, fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Header (model.BlockHeader) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return bytesRead, err
		}
		bytesRead++
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return bytesRead, err
			}
			bytesRead--
			t.Header = new(BlockHeader)
			if read, err := t.Header.UnmarshalCBOR(br); err != nil {
				return bytesRead, xerrors.Errorf("unmarshaling t.Header pointer: %w", err)
			} else {
				bytesRead += read
			}
		}

	}
	// t.Txs (model.TransactionSlice) (slice)

	maj, extra, read, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return bytesRead, err
	}
	bytesRead += read

	if extra > cbg.MaxLength {
		return bytesRead, fmt.Errorf("t.Txs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return bytesRead, fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Txs = make([]Transaction, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v Transaction
		if read, err := v.UnmarshalCBOR(br); err != nil {
			return bytesRead, err
		} else {
			bytesRead += read
		}

		t.Txs[i] = v
	}

	return bytesRead, nil
}