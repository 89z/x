// Go package for low-level interaction with the operating system
package sys
//go:generate mkwinsyscall -output zsys.go sys.go
//sys MsiDatabaseOpenView(database int, query string, view *int) (e error) = msi.MsiDatabaseOpenViewW
//sys MsiInstallProduct(path string, command string) (e error) = msi.MsiInstallProductW
//sys MsiOpenDatabase(path string, persist int, database *int) (e error) = msi.MsiOpenDatabaseW
//sys MsiRecordGetString(record int, field int, buf *uint16, bufSize *int) (e error) = msi.MsiRecordGetStringW
//sys MsiViewExecute(view int, record int) (e error) = msi.MsiViewExecute
//sys MsiViewFetch(view int, record *int) (e error) = msi.MsiViewFetch
//sys ShellExecute(hwnd int, oper string, file string, param string, dir string, show int) (err error) = shell32.ShellExecuteW
