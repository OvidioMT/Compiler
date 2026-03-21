@aquiStr = global [2 x i8] c"1\00"

define i32 @main() {
0:
	%1 = alloca i32
	store i32 0, i32* %1
	%2 = load i32, i32* %1
	%3 = load i32, i32* %1
	%4 = load i32, i32* %1
	%5 = call i32 @puts([2 x i8]* @aquiStr)
	ret i32 0
}

declare i32 @puts(i8* %0)
