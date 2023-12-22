use std::mem::MaybeUninit;
use std::slice;

#[no_mangle]
pub unsafe extern "C" fn required_gas(input_ptr: *mut u8, input_len: u32) -> u64 {
    let slice = slice::from_raw_parts_mut(input_ptr, input_len as usize);
    return slice.len() as u64
}

#[no_mangle]
pub unsafe extern "C" fn run(input_ptr: *mut u8, input_len: u32) ->u64 {
    return ((input_ptr as u64) << 32) | input_len as u64
}


#[no_mangle]
pub extern "C" fn add(x: u32, y: u32) -> u32 {
    return x+y
}

#[cfg_attr(all(target_arch = "wasm32"), export_name = "allocate")]
#[no_mangle]
pub extern "C" fn _allocate(size: u32) -> *mut u8 {
    allocate(size as usize)
}

/// Allocates size bytes and leaks the pointer where they start.
fn allocate(size: usize) -> *mut u8 {
    // Allocate the amount of bytes needed.
    let vec: Vec<MaybeUninit<u8>> = Vec::with_capacity(size);

    // into_raw leaks the memory to the caller.
    Box::into_raw(vec.into_boxed_slice()) as *mut u8
}

/// WebAssembly export that deallocates a pointer of the given size (linear
/// memory offset, byteCount) allocated by [`allocate`].
#[cfg_attr(all(target_arch = "wasm32"), export_name = "deallocate")]
#[no_mangle]
pub unsafe extern "C" fn _deallocate(ptr: u32, size: u32) {
    deallocate(ptr as *mut u8, size as usize);
}

/// Retakes the pointer which allows its memory to be freed.
unsafe fn deallocate(ptr: *mut u8, size: usize) {
    let _ = Vec::from_raw_parts(ptr, 0, size);
}
