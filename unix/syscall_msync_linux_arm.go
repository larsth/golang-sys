package unix

//Msync is the implementation of the musl v1.14 libc msync function:
//See http://git.musl-libc.org/cgit/musl/tree/src/mman/msync.c
//int msync(void *start, size_t len, int flags)
//{
//	return syscall_cp(SYS_msync, start, len, flags);
//}
func Msync(data []byte, len ) {

}
