package precompiles

/*
#cgo LDFLAGS: -L../target/release/ -lprecompiles -lm

#include <stdbool.h>
#include <stdint.h>

uint64_t __precompile_anonymous_verify_gas(const void* data_ptr, const uint32_t data_len);
uint8_t __precompile_anonymous_verify(const void* data_ptr, const uint32_t data_len);

uint64_t __precompile_anemoi_gas(const void* data_ptr, const uint32_t data_len);
uint8_t __precompile_anemoi(const void* data_ptr, const uint32_t data_len, void* ret_val);

uint64_t __precompile_shuffle_verify_gas(const void* data_ptr, const uint32_t data_len);
uint8_t __precompile_shuffle_verify(const void* data_ptr, const uint32_t data_len);

uint64_t __precompile_shuffle_exec_gas(const void* data_ptr, const uint32_t data_len);
uint8_t __precompile_shuffle_exec(const void* data_ptr, const uint32_t data_len, void* ret_val, void* out_len);

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Anonymous struct{}

func (a *Anonymous) RequiredGas(input []byte) uint64 {
	cstr := unsafe.Pointer(&input[0])
	len := C.uint(len(input))

	gas := C.__precompile_anonymous_verify_gas(cstr, len)

	return uint64(gas)
}

func (a *Anonymous) Run(input []byte) ([]byte, error) {
	cstr := unsafe.Pointer(&input[0])
	len := C.uint(len(input))

	res := C.__precompile_anonymous_verify(cstr, len)

	output := make([]byte, 32)

	output[31] = byte(res)

	return output, nil
}

type Anemoi struct{}

func (a *Anemoi) RequiredGas(input []byte) uint64 {
	cstr := unsafe.Pointer(&input[0])
	len := C.uint(len(input))

	gas := C.__precompile_anemoi_gas(cstr, len)

	return uint64(gas)
}

func (a *Anemoi) Run(input []byte) ([]byte, error) {
	output := make([]byte, 64)
	cout := unsafe.Pointer(&output[0])

	cstr := unsafe.Pointer(&input[0])
	len := C.uint(len(input))

	res := C.__precompile_anemoi(cstr, len, cout)

	output[63] = byte(res)

	return output, nil
}

type ShuffleVerify struct{}

func (a *ShuffleVerify) RequiredGas(input []byte) uint64 {
	cstr := unsafe.Pointer(&input[0])
	len := C.uint(len(input))

	gas := C.__precompile_shuffle_verify_gas(cstr, len)

	return uint64(gas)
}

func (a *ShuffleVerify) Run(input []byte) ([]byte, error) {
	cstr := unsafe.Pointer(&input[0])
	len := C.uint(len(input))

	res := C.__precompile_shuffle_verify(cstr, len)

	output := make([]byte, 32)

	output[31] = byte(res)

	return output, nil
}

type ShuffleExec struct{}

func (a *ShuffleExec) RequiredGas(input []byte) uint64 {
	cstr := unsafe.Pointer(&input[0])
	len := C.uint(len(input))

	gas := C.__precompile_shuffle_exec_gas(cstr, len)

	return uint64(gas)
}

func (a *ShuffleExec) Run(input []byte) ([]byte, error) {
	output := make([]byte, 1024)
	cout := unsafe.Pointer(&output[0])

	cstr := unsafe.Pointer(&input[0])
	len := C.uint(len(input))
	out_len := C.uint(0)
	out_len_ptr := unsafe.Pointer(&out_len)

	res := C.__precompile_shuffle_exec(cstr, len, cout, out_len_ptr)

	output[1023] = byte(res)

	output2 := make([]byte, out_len)
	copy(output2, output[:out_len])

	var err error = nil
	if res != 0 {
		err = fmt.Errorf("error: %d", res)
	}

	return output2[:], err
}
