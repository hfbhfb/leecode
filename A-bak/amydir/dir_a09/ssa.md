
## go tool compile -S ssa_01.go 


``` s
"".Helloworld STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0
        0x0000 00000 (ssa.go:3) TEXT    "".Helloworld(SB), NOSPLIT|ABIInternal, $0-0
        0x0000 00000 (ssa.go:3) FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (ssa.go:3) FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (ssa.go:5) RET
        0x0000 c3                                               .
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 64 69 72 61 30 39                                dira09
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
```


##  go tool compile -S ssa_02.go 

``` s
"".Y02 STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0
        0x0000 00000 (ssa_02.go:3)      TEXT    "".Y02(SB), NOSPLIT|ABIInternal, $0-0
        0x0000 00000 (ssa_02.go:3)      FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (ssa_02.go:3)      FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (ssa_02.go:6)      RET
        0x0000 c3                                               .
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 64 69 72 61 30 39                                dira09
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
```


## go tool compile -S ssa_03.go 

``` s
"".Z03 STEXT nosplit size=1 args=0x0 locals=0x0 funcid=0x0
        0x0000 00000 (ssa_03.go:3)      TEXT    "".Z03(SB), NOSPLIT|ABIInternal, $0-0
        0x0000 00000 (ssa_03.go:3)      FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (ssa_03.go:3)      FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (ssa_03.go:6)      RET
        0x0000 c3                                               .
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 64 69 72 61 30 39                                dira09
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
hfbdeMBP:dir_a09 hfb$ 
hfbdeMBP:dir_a09 hfb$ 
hfbdeMBP:dir_a09 hfb$ go tool compile -S ssa_03.go 
"".Z03 STEXT size=103 args=0x0 locals=0x48 funcid=0x0
        0x0000 00000 (ssa_03.go:3)      TEXT    "".Z03(SB), ABIInternal, $72-0
        0x0000 00000 (ssa_03.go:3)      MOVQ    (TLS), CX
        0x0009 00009 (ssa_03.go:3)      CMPQ    SP, 16(CX)
        0x000d 00013 (ssa_03.go:3)      PCDATA  $0, $-2
        0x000d 00013 (ssa_03.go:3)      JLS     96
        0x000f 00015 (ssa_03.go:3)      PCDATA  $0, $-1
        0x000f 00015 (ssa_03.go:3)      SUBQ    $72, SP
        0x0013 00019 (ssa_03.go:3)      MOVQ    BP, 64(SP)
        0x0018 00024 (ssa_03.go:3)      LEAQ    64(SP), BP
        0x001d 00029 (ssa_03.go:3)      FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (ssa_03.go:3)      FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (ssa_03.go:5)      LEAQ    type.int(SB), AX
        0x0024 00036 (ssa_03.go:5)      MOVQ    AX, (SP)
        0x0028 00040 (ssa_03.go:5)      LEAQ    ""..autotmp_4+64(SP), AX
        0x002d 00045 (ssa_03.go:5)      MOVQ    AX, 8(SP)
        0x0032 00050 (ssa_03.go:5)      XORPS   X0, X0
        0x0035 00053 (ssa_03.go:5)      MOVUPS  X0, 16(SP)
        0x003a 00058 (ssa_03.go:5)      MOVQ    $1, 32(SP)
        0x0043 00067 (ssa_03.go:5)      PCDATA  $1, $0
        0x0043 00067 (ssa_03.go:5)      CALL    runtime.growslice(SB)
        0x0048 00072 (ssa_03.go:5)      MOVQ    40(SP), AX
        0x004d 00077 (ssa_03.go:5)      MOVQ    $3, (AX)
        0x0054 00084 (ssa_03.go:8)      MOVQ    64(SP), BP
        0x0059 00089 (ssa_03.go:8)      ADDQ    $72, SP
        0x005d 00093 (ssa_03.go:8)      RET
        0x005e 00094 (ssa_03.go:8)      NOP
        0x005e 00094 (ssa_03.go:3)      PCDATA  $1, $-1
        0x005e 00094 (ssa_03.go:3)      PCDATA  $0, $-2
        0x005e 00094 (ssa_03.go:3)      NOP
        0x0060 00096 (ssa_03.go:3)      CALL    runtime.morestack_noctxt(SB)
        0x0065 00101 (ssa_03.go:3)      PCDATA  $0, $-1
        0x0065 00101 (ssa_03.go:3)      JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 51 48  eH..%....H;a.vQH
        0x0010 83 ec 48 48 89 6c 24 40 48 8d 6c 24 40 48 8d 05  ..HH.l$@H.l$@H..
        0x0020 00 00 00 00 48 89 04 24 48 8d 44 24 40 48 89 44  ....H..$H.D$@H.D
        0x0030 24 08 0f 57 c0 0f 11 44 24 10 48 c7 44 24 20 01  $..W...D$.H.D$ .
        0x0040 00 00 00 e8 00 00 00 00 48 8b 44 24 28 48 c7 00  ........H.D$(H..
        0x0050 03 00 00 00 48 8b 6c 24 40 48 83 c4 48 c3 66 90  ....H.l$@H..H.f.
        0x0060 e8 00 00 00 00 eb 99                             .......
        rel 5+4 t=17 TLS+0
        rel 32+4 t=16 type.int+0
        rel 68+4 t=8 runtime.growslice+0
        rel 97+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 64 69 72 61 30 39                                dira09
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
```



##  go tool compile -S ssa_04.go 
``` s
"".M04 STEXT size=138 args=0x18 locals=0x48 funcid=0x0
        0x0000 00000 (ssa_04.go:3)      TEXT    "".M04(SB), ABIInternal, $72-24
        0x0000 00000 (ssa_04.go:3)      MOVQ    (TLS), CX
        0x0009 00009 (ssa_04.go:3)      CMPQ    SP, 16(CX)
        0x000d 00013 (ssa_04.go:3)      PCDATA  $0, $-2
        0x000d 00013 (ssa_04.go:3)      JLS     128
        0x000f 00015 (ssa_04.go:3)      PCDATA  $0, $-1
        0x000f 00015 (ssa_04.go:3)      SUBQ    $72, SP
        0x0013 00019 (ssa_04.go:3)      MOVQ    BP, 64(SP)
        0x0018 00024 (ssa_04.go:3)      LEAQ    64(SP), BP
        0x001d 00029 (ssa_04.go:3)      FUNCDATA        $0, gclocals·2a5305abe05176240e61b8620e19a815(SB)
        0x001d 00029 (ssa_04.go:3)      FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (ssa_04.go:5)      LEAQ    type.int(SB), AX
        0x0024 00036 (ssa_04.go:5)      MOVQ    AX, (SP)
        0x0028 00040 (ssa_04.go:5)      LEAQ    runtime.zerobase(SB), AX
        0x002f 00047 (ssa_04.go:5)      MOVQ    AX, 8(SP)
        0x0034 00052 (ssa_04.go:5)      XORPS   X0, X0
        0x0037 00055 (ssa_04.go:5)      MOVUPS  X0, 16(SP)
        0x003c 00060 (ssa_04.go:5)      MOVQ    $1, 32(SP)
        0x0045 00069 (ssa_04.go:5)      PCDATA  $1, $0
        0x0045 00069 (ssa_04.go:5)      CALL    runtime.growslice(SB)
        0x004a 00074 (ssa_04.go:5)      MOVQ    40(SP), AX
        0x004f 00079 (ssa_04.go:5)      MOVQ    48(SP), CX
        0x0054 00084 (ssa_04.go:5)      MOVQ    56(SP), DX
        0x0059 00089 (ssa_04.go:5)      MOVQ    $3, (AX)
        0x0060 00096 (ssa_04.go:7)      MOVQ    AX, "".~r0+80(SP)
        0x0065 00101 (ssa_04.go:5)      LEAQ    1(CX), AX
        0x0069 00105 (ssa_04.go:7)      MOVQ    AX, "".~r0+88(SP)
        0x006e 00110 (ssa_04.go:7)      MOVQ    DX, "".~r0+96(SP)
        0x0073 00115 (ssa_04.go:7)      MOVQ    64(SP), BP
        0x0078 00120 (ssa_04.go:7)      ADDQ    $72, SP
        0x007c 00124 (ssa_04.go:7)      RET
        0x007d 00125 (ssa_04.go:7)      NOP
        0x007d 00125 (ssa_04.go:3)      PCDATA  $1, $-1
        0x007d 00125 (ssa_04.go:3)      PCDATA  $0, $-2
        0x007d 00125 (ssa_04.go:3)      NOP
        0x0080 00128 (ssa_04.go:3)      CALL    runtime.morestack_noctxt(SB)
        0x0085 00133 (ssa_04.go:3)      PCDATA  $0, $-1
        0x0085 00133 (ssa_04.go:3)      JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 71 48  eH..%....H;a.vqH
        0x0010 83 ec 48 48 89 6c 24 40 48 8d 6c 24 40 48 8d 05  ..HH.l$@H.l$@H..
        0x0020 00 00 00 00 48 89 04 24 48 8d 05 00 00 00 00 48  ....H..$H......H
        0x0030 89 44 24 08 0f 57 c0 0f 11 44 24 10 48 c7 44 24  .D$..W...D$.H.D$
        0x0040 20 01 00 00 00 e8 00 00 00 00 48 8b 44 24 28 48   .........H.D$(H
        0x0050 8b 4c 24 30 48 8b 54 24 38 48 c7 00 03 00 00 00  .L$0H.T$8H......
        0x0060 48 89 44 24 50 48 8d 41 01 48 89 44 24 58 48 89  H.D$PH.A.H.D$XH.
        0x0070 54 24 60 48 8b 6c 24 40 48 83 c4 48 c3 0f 1f 00  T$`H.l$@H..H....
        0x0080 e8 00 00 00 00 e9 76 ff ff ff                    ......v...
        rel 5+4 t=17 TLS+0
        rel 32+4 t=16 type.int+0
        rel 43+4 t=16 runtime.zerobase+0
        rel 70+4 t=8 runtime.growslice+0
        rel 129+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 64 69 72 61 30 39                                dira09
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*[]int- SRODATA dupok size=9
        0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 1b 31 52 88 08 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]int-+0
        rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]int-+0
        rel 44+4 t=6 type.*[]int+0
        rel 48+8 t=1 type.int+0
runtime.memequal0·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal0+0
type..namedata.*[0]int- SRODATA dupok size=10
        0x0000 00 00 07 2a 5b 30 5d 69 6e 74                    ...*[0]int
type.*[0]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 f4 3c 85 c0 08 08 08 36 00 00 00 00 00 00 00 00  .<.....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[0]int-+0
        rel 48+8 t=1 type.[0]int+0
runtime.gcbits. SRODATA dupok size=0
type.[0]int SRODATA dupok size=72
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 c1 04 70 f1 0a 08 08 11 00 00 00 00 00 00 00 00  ..p.............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal0·f+0
        rel 32+8 t=1 runtime.gcbits.+0
        rel 40+4 t=5 type..namedata.*[0]int-+0
        rel 44+4 t=6 type.*[0]int+0
        rel 48+8 t=1 type.int+0
        rel 56+8 t=1 type.[]int+0
gclocals·2a5305abe05176240e61b8620e19a815 SRODATA dupok size=9
        0x0000 01 00 00 00 01 00 00 00 00                       .........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
```






## go tool compile -S ssa_05.go 
``` s
"".M05 STEXT size=234 args=0x18 locals=0x28 funcid=0x0
        0x0000 00000 (ssa_05.go:3)      TEXT    "".M05(SB), ABIInternal, $40-24
        0x0000 00000 (ssa_05.go:3)      MOVQ    (TLS), CX
        0x0009 00009 (ssa_05.go:3)      CMPQ    SP, 16(CX)
        0x000d 00013 (ssa_05.go:3)      PCDATA  $0, $-2
        0x000d 00013 (ssa_05.go:3)      JLS     224
        0x0013 00019 (ssa_05.go:3)      PCDATA  $0, $-1
        0x0013 00019 (ssa_05.go:3)      SUBQ    $40, SP
        0x0017 00023 (ssa_05.go:3)      MOVQ    BP, 32(SP)
        0x001c 00028 (ssa_05.go:3)      LEAQ    32(SP), BP
        0x0021 00033 (ssa_05.go:3)      FUNCDATA        $0, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
        0x0021 00033 (ssa_05.go:3)      FUNCDATA        $1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0021 00033 (ssa_05.go:4)      LEAQ    type.[12]int(SB), AX
        0x0028 00040 (ssa_05.go:4)      MOVQ    AX, (SP)
        0x002c 00044 (ssa_05.go:4)      PCDATA  $1, $0
        0x002c 00044 (ssa_05.go:4)      CALL    runtime.newobject(SB)
        0x0031 00049 (ssa_05.go:4)      MOVQ    8(SP), AX
        0x0036 00054 (ssa_05.go:4)      MOVQ    AX, ""..autotmp_18+24(SP)
        0x003b 00059 (ssa_05.go:4)      MOVQ    $3, (AX)
        0x0042 00066 (ssa_05.go:4)      MOVQ    $2, 8(AX)
        0x004a 00074 (ssa_05.go:4)      MOVQ    $8, 16(AX)
        0x0052 00082 (ssa_05.go:4)      MOVQ    $9, 24(AX)
        0x005a 00090 (ssa_05.go:4)      MOVQ    $9, 32(AX)
        0x0062 00098 (ssa_05.go:4)      MOVQ    $9, 40(AX)
        0x006a 00106 (ssa_05.go:4)      MOVQ    $9, 48(AX)
        0x0072 00114 (ssa_05.go:4)      MOVQ    $9, 56(AX)
        0x007a 00122 (ssa_05.go:4)      MOVQ    $9, 64(AX)
        0x0082 00130 (ssa_05.go:4)      MOVQ    $9, 72(AX)
        0x008a 00138 (ssa_05.go:4)      MOVQ    $9, 80(AX)
        0x0092 00146 (ssa_05.go:4)      MOVQ    $9, 88(AX)
        0x009a 00154 (ssa_05.go:6)      LEAQ    40(AX), CX
        0x009e 00158 (ssa_05.go:6)      MOVQ    CX, (SP)
        0x00a2 00162 (ssa_05.go:6)      LEAQ    48(AX), CX
        0x00a6 00166 (ssa_05.go:6)      MOVQ    CX, 8(SP)
        0x00ab 00171 (ssa_05.go:6)      MOVQ    $48, 16(SP)
        0x00b4 00180 (ssa_05.go:6)      PCDATA  $1, $1
        0x00b4 00180 (ssa_05.go:6)      CALL    runtime.memmove(SB)
        0x00b9 00185 (ssa_05.go:7)      MOVQ    ""..autotmp_18+24(SP), AX
        0x00be 00190 (ssa_05.go:7)      MOVQ    AX, "".~r0+48(SP)
        0x00c3 00195 (ssa_05.go:7)      MOVQ    $11, "".~r0+56(SP)
        0x00cc 00204 (ssa_05.go:7)      MOVQ    $12, "".~r0+64(SP)
        0x00d5 00213 (ssa_05.go:7)      MOVQ    32(SP), BP
        0x00da 00218 (ssa_05.go:7)      ADDQ    $40, SP
        0x00de 00222 (ssa_05.go:7)      RET
        0x00df 00223 (ssa_05.go:7)      NOP
        0x00df 00223 (ssa_05.go:3)      PCDATA  $1, $-1
        0x00df 00223 (ssa_05.go:3)      PCDATA  $0, $-2
        0x00df 00223 (ssa_05.go:3)      NOP
        0x00e0 00224 (ssa_05.go:3)      CALL    runtime.morestack_noctxt(SB)
        0x00e5 00229 (ssa_05.go:3)      PCDATA  $0, $-1
        0x00e5 00229 (ssa_05.go:3)      JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 cd  eH..%....H;a....
        0x0010 00 00 00 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24  ...H..(H.l$ H.l$
        0x0020 20 48 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00   H......H..$....
        0x0030 00 48 8b 44 24 08 48 89 44 24 18 48 c7 00 03 00  .H.D$.H.D$.H....
        0x0040 00 00 48 c7 40 08 02 00 00 00 48 c7 40 10 08 00  ..H.@.....H.@...
        0x0050 00 00 48 c7 40 18 09 00 00 00 48 c7 40 20 09 00  ..H.@.....H.@ ..
        0x0060 00 00 48 c7 40 28 09 00 00 00 48 c7 40 30 09 00  ..H.@(....H.@0..
        0x0070 00 00 48 c7 40 38 09 00 00 00 48 c7 40 40 09 00  ..H.@8....H.@@..
        0x0080 00 00 48 c7 40 48 09 00 00 00 48 c7 40 50 09 00  ..H.@H....H.@P..
        0x0090 00 00 48 c7 40 58 09 00 00 00 48 8d 48 28 48 89  ..H.@X....H.H(H.
        0x00a0 0c 24 48 8d 48 30 48 89 4c 24 08 48 c7 44 24 10  .$H.H0H.L$.H.D$.
        0x00b0 30 00 00 00 e8 00 00 00 00 48 8b 44 24 18 48 89  0........H.D$.H.
        0x00c0 44 24 30 48 c7 44 24 38 0b 00 00 00 48 c7 44 24  D$0H.D$8....H.D$
        0x00d0 40 0c 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 90  @....H.l$ H..(..
        0x00e0 e8 00 00 00 00 e9 16 ff ff ff                    ..........
        rel 5+4 t=17 TLS+0
        rel 36+4 t=16 type.[12]int+0
        rel 45+4 t=8 runtime.newobject+0
        rel 181+4 t=8 runtime.memmove+0
        rel 225+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 64 69 72 61 30 39                                dira09
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*[]int- SRODATA dupok size=9
        0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 1b 31 52 88 08 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]int-+0
        rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]int-+0
        rel 44+4 t=6 type.*[]int+0
        rel 48+8 t=1 type.int+0
type..eqfunc96 SRODATA dupok size=16
        0x0000 00 00 00 00 00 00 00 00 60 00 00 00 00 00 00 00  ........`.......
        rel 0+8 t=1 runtime.memequal_varlen+0
type..namedata.*[12]int- SRODATA dupok size=11
        0x0000 00 00 08 2a 5b 31 32 5d 69 6e 74                 ...*[12]int
type.*[12]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 e3 d3 e0 ea 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[12]int-+0
        rel 48+8 t=1 type.[12]int+0
runtime.gcbits. SRODATA dupok size=0
type.[12]int SRODATA dupok size=72
        0x0000 60 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  `...............
        0x0010 90 af 57 27 0a 08 08 11 00 00 00 00 00 00 00 00  ..W'............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 0c 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 type..eqfunc96+0
        rel 32+8 t=1 runtime.gcbits.+0
        rel 40+4 t=5 type..namedata.*[12]int-+0
        rel 44+4 t=6 type.*[12]int+0
        rel 48+8 t=1 type.int+0
        rel 56+8 t=1 type.[]int+0
gclocals·263043c8f03e3241528dfae4e2812ef4 SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 01                    ..........

```

##  go tool compile -S ssa_06.go 
``` s
"".M06 STEXT size=248 args=0x18 locals=0x130 funcid=0x0
        0x0000 00000 (ssa_06.go:3)      TEXT    "".M06(SB), ABIInternal, $304-24
        0x0000 00000 (ssa_06.go:3)      MOVQ    (TLS), CX
        0x0009 00009 (ssa_06.go:3)      LEAQ    -176(SP), AX
        0x0011 00017 (ssa_06.go:3)      CMPQ    AX, 16(CX)
        0x0015 00021 (ssa_06.go:3)      PCDATA  $0, $-2
        0x0015 00021 (ssa_06.go:3)      JLS     238
        0x001b 00027 (ssa_06.go:3)      PCDATA  $0, $-1
        0x001b 00027 (ssa_06.go:3)      SUBQ    $304, SP
        0x0022 00034 (ssa_06.go:3)      MOVQ    BP, 296(SP)
        0x002a 00042 (ssa_06.go:3)      LEAQ    296(SP), BP
        0x0032 00050 (ssa_06.go:3)      FUNCDATA        $0, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
        0x0032 00050 (ssa_06.go:3)      FUNCDATA        $1, gclocals·6d8d66cc8b183c896811c0201ffd5671(SB)
        0x0032 00050 (ssa_06.go:3)      FUNCDATA        $2, "".M06.stkobj(SB)
        0x0032 00050 (ssa_06.go:4)      XORPS   X0, X0
        0x0035 00053 (ssa_06.go:4)      MOVUPS  X0, ""..autotmp_2+40(SP)
        0x003a 00058 (ssa_06.go:4)      MOVUPS  X0, ""..autotmp_2+56(SP)
        0x003f 00063 (ssa_06.go:4)      MOVUPS  X0, ""..autotmp_2+72(SP)
        0x0044 00068 (ssa_06.go:4)      LEAQ    ""..autotmp_3+88(SP), DI
        0x0049 00073 (ssa_06.go:4)      PCDATA  $0, $-2
        0x0049 00073 (ssa_06.go:4)      LEAQ    -48(DI), DI
        0x004d 00077 (ssa_06.go:4)      NOP
        0x0060 00096 (ssa_06.go:4)      DUFFZERO        $239
        0x0073 00115 (ssa_06.go:4)      PCDATA  $0, $-1
        0x0073 00115 (ssa_06.go:4)      LEAQ    ""..autotmp_3+88(SP), AX
        0x0078 00120 (ssa_06.go:4)      MOVQ    AX, ""..autotmp_2+56(SP)
        0x007d 00125 (ssa_06.go:4)      PCDATA  $1, $1
        0x007d 00125 (ssa_06.go:4)      NOP
        0x0080 00128 (ssa_06.go:4)      CALL    runtime.fastrand(SB)
        0x0085 00133 (ssa_06.go:4)      MOVL    (SP), AX
        0x0088 00136 (ssa_06.go:4)      MOVL    AX, ""..autotmp_2+52(SP)
        0x008c 00140 (ssa_06.go:6)      LEAQ    type.map[string]int(SB), AX
        0x0093 00147 (ssa_06.go:6)      MOVQ    AX, (SP)
        0x0097 00151 (ssa_06.go:6)      LEAQ    ""..autotmp_2+40(SP), AX
        0x009c 00156 (ssa_06.go:6)      MOVQ    AX, 8(SP)
        0x00a1 00161 (ssa_06.go:6)      LEAQ    go.string."aa"(SB), AX
        0x00a8 00168 (ssa_06.go:6)      MOVQ    AX, 16(SP)
        0x00ad 00173 (ssa_06.go:6)      MOVQ    $2, 24(SP)
        0x00b6 00182 (ssa_06.go:6)      PCDATA  $1, $0
        0x00b6 00182 (ssa_06.go:6)      CALL    runtime.mapassign_faststr(SB)
        0x00bb 00187 (ssa_06.go:6)      MOVQ    32(SP), AX
        0x00c0 00192 (ssa_06.go:6)      MOVQ    $8, (AX)
        0x00c7 00199 (ssa_06.go:8)      MOVQ    $0, "".~r0+312(SP)
        0x00d3 00211 (ssa_06.go:8)      XORPS   X1, X1
        0x00d6 00214 (ssa_06.go:8)      MOVUPS  X1, "".~r0+320(SP)
        0x00de 00222 (ssa_06.go:8)      MOVQ    296(SP), BP
        0x00e6 00230 (ssa_06.go:8)      ADDQ    $304, SP
        0x00ed 00237 (ssa_06.go:8)      RET
        0x00ee 00238 (ssa_06.go:8)      NOP
        0x00ee 00238 (ssa_06.go:3)      PCDATA  $1, $-1
        0x00ee 00238 (ssa_06.go:3)      PCDATA  $0, $-2
        0x00ee 00238 (ssa_06.go:3)      CALL    runtime.morestack_noctxt(SB)
        0x00f3 00243 (ssa_06.go:3)      PCDATA  $0, $-1
        0x00f3 00243 (ssa_06.go:3)      JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 84 24 50 ff ff  eH..%....H..$P..
        0x0010 ff 48 3b 41 10 0f 86 d3 00 00 00 48 81 ec 30 01  .H;A.......H..0.
        0x0020 00 00 48 89 ac 24 28 01 00 00 48 8d ac 24 28 01  ..H..$(...H..$(.
        0x0030 00 00 0f 57 c0 0f 11 44 24 28 0f 11 44 24 38 0f  ...W...D$(..D$8.
        0x0040 11 44 24 48 48 8d 7c 24 58 48 8d 7f d0 66 0f 1f  .D$HH.|$XH...f..
        0x0050 84 00 00 00 00 00 66 0f 1f 84 00 00 00 00 00 90  ......f.........
        0x0060 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00 00 00 48  H.l$.H.l$......H
        0x0070 8b 6d 00 48 8d 44 24 58 48 89 44 24 38 0f 1f 00  .m.H.D$XH.D$8...
        0x0080 e8 00 00 00 00 8b 04 24 89 44 24 34 48 8d 05 00  .......$.D$4H...
        0x0090 00 00 00 48 89 04 24 48 8d 44 24 28 48 89 44 24  ...H..$H.D$(H.D$
        0x00a0 08 48 8d 05 00 00 00 00 48 89 44 24 10 48 c7 44  .H......H.D$.H.D
        0x00b0 24 18 02 00 00 00 e8 00 00 00 00 48 8b 44 24 20  $..........H.D$ 
        0x00c0 48 c7 00 08 00 00 00 48 c7 84 24 38 01 00 00 00  H......H..$8....
        0x00d0 00 00 00 0f 57 c9 0f 11 8c 24 40 01 00 00 48 8b  ....W....$@...H.
        0x00e0 ac 24 28 01 00 00 48 81 c4 30 01 00 00 c3 e8 00  .$(...H..0......
        0x00f0 00 00 00 e9 08 ff ff ff                          ........
        rel 5+4 t=17 TLS+0
        rel 107+4 t=8 runtime.duffzero+239
        rel 129+4 t=8 runtime.fastrand+0
        rel 143+4 t=16 type.map[string]int+0
        rel 164+4 t=16 go.string."aa"+0
        rel 183+4 t=8 runtime.mapassign_faststr+0
        rel 239+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 64 69 72 61 30 39                                dira09
go.string."aa" SRODATA dupok size=2
        0x0000 61 61                                            aa
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*[]uint8- SRODATA dupok size=11
        0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 a5 8e d0 69 08 08 08 36 00 00 00 00 00 00 00 00  ...i...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]uint8-+0
        rel 48+8 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 df 7e 2e 38 02 08 08 17 00 00 00 00 00 00 00 00  .~.8............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]uint8-+0
        rel 44+4 t=6 type.*[]uint8+0
        rel 48+8 t=1 type.uint8+0
type..namedata.*[8]uint8- SRODATA dupok size=12
        0x0000 00 00 09 2a 5b 38 5d 75 69 6e 74 38              ...*[8]uint8
type.*[8]uint8 SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 a9 89 a5 7a 08 08 08 36 00 00 00 00 00 00 00 00  ...z...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[8]uint8-+0
        rel 48+8 t=1 type.[8]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[8]uint8 SRODATA dupok size=72
        0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 3e f9 30 b4 0a 01 01 11 00 00 00 00 00 00 00 00  >.0.............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 08 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.+0
        rel 40+4 t=5 type..namedata.*[8]uint8-+0
        rel 44+4 t=6 type.*[8]uint8+0
        rel 48+8 t=1 type.uint8+0
        rel 56+8 t=1 type.[]uint8+0
type..namedata.*[]string- SRODATA dupok size=12
        0x0000 00 00 09 2a 5b 5d 73 74 72 69 6e 67              ...*[]string
type.*[]string SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 92 22 76 84 08 08 08 36 00 00 00 00 00 00 00 00  ."v....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]string-+0
        rel 48+8 t=1 type.[]string+0
type.[]string SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 d3 a8 f3 0a 02 08 08 17 00 00 00 00 00 00 00 00  ................
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]string-+0
        rel 44+4 t=6 type.*[]string+0
        rel 48+8 t=1 type.string+0
type..namedata.*[8]string- SRODATA dupok size=13
        0x0000 00 00 0a 2a 5b 38 5d 73 74 72 69 6e 67           ...*[8]string
type.*[8]string SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 ad 94 14 6f 08 08 08 36 00 00 00 00 00 00 00 00  ...o...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[8]string-+0
        rel 48+8 t=1 type.noalg.[8]string+0
runtime.gcbits.5555 SRODATA dupok size=2
        0x0000 55 55                                            UU
type.noalg.[8]string SRODATA dupok size=72
        0x0000 80 00 00 00 00 00 00 00 78 00 00 00 00 00 00 00  ........x.......
        0x0010 55 53 8c 3e 02 08 08 11 00 00 00 00 00 00 00 00  US.>............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 08 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.5555+0
        rel 40+4 t=5 type..namedata.*[8]string-+0
        rel 44+4 t=6 type.*[8]string+0
        rel 48+8 t=1 type.string+0
        rel 56+8 t=1 type.[]string+0
type..namedata.*[]int- SRODATA dupok size=9
        0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 1b 31 52 88 08 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]int-+0
        rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]int-+0
        rel 44+4 t=6 type.*[]int+0
        rel 48+8 t=1 type.int+0
type..namedata.*[8]int- SRODATA dupok size=10
        0x0000 00 00 07 2a 5b 38 5d 69 6e 74                    ...*[8]int
type.*[8]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 f3 3f a8 3b 08 08 08 36 00 00 00 00 00 00 00 00  .?.;...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[8]int-+0
        rel 48+8 t=1 type.noalg.[8]int+0
type.noalg.[8]int SRODATA dupok size=72
        0x0000 40 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  @...............
        0x0010 96 99 d5 05 02 08 08 11 00 00 00 00 00 00 00 00  ................
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 08 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.+0
        rel 40+4 t=5 type..namedata.*[8]int-+0
        rel 44+4 t=6 type.*[8]int+0
        rel 48+8 t=1 type.int+0
        rel 56+8 t=1 type.[]int+0
runtime.gcbits.aaaa0002 SRODATA dupok size=4
        0x0000 aa aa 00 02                                      ....
type..namedata.*map.bucket[string]int- SRODATA dupok size=25
        0x0000 00 00 16 2a 6d 61 70 2e 62 75 63 6b 65 74 5b 73  ...*map.bucket[s
        0x0010 74 72 69 6e 67 5d 69 6e 74                       tring]int
type..importpath.. SRODATA dupok size=3
        0x0000 00 00 00                                         ...
type..namedata.topbits- SRODATA dupok size=10
        0x0000 00 00 07 74 6f 70 62 69 74 73                    ...topbits
type..namedata.keys- SRODATA dupok size=7
        0x0000 00 00 04 6b 65 79 73                             ...keys
type..namedata.elems- SRODATA dupok size=8
        0x0000 00 00 05 65 6c 65 6d 73                          ...elems
type..namedata.overflow- SRODATA dupok size=11
        0x0000 00 00 08 6f 76 65 72 66 6c 6f 77                 ...overflow
type.noalg.map.bucket[string]int SRODATA dupok size=176
        0x0000 d0 00 00 00 00 00 00 00 d0 00 00 00 00 00 00 00  ................
        0x0010 5d 68 63 71 02 08 08 19 00 00 00 00 00 00 00 00  ]hcq............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 04 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0090 10 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x00a0 00 00 00 00 00 00 00 00 90 01 00 00 00 00 00 00  ................
        rel 32+8 t=1 runtime.gcbits.aaaa0002+0
        rel 40+4 t=5 type..namedata.*map.bucket[string]int-+0
        rel 44+4 t=6 type.*map.bucket[string]int+0
        rel 48+8 t=1 type..importpath..+0
        rel 56+8 t=1 type.noalg.map.bucket[string]int+80
        rel 80+8 t=1 type..namedata.topbits-+0
        rel 88+8 t=1 type.[8]uint8+0
        rel 104+8 t=1 type..namedata.keys-+0
        rel 112+8 t=1 type.noalg.[8]string+0
        rel 128+8 t=1 type..namedata.elems-+0
        rel 136+8 t=1 type.noalg.[8]int+0
        rel 152+8 t=1 type..namedata.overflow-+0
        rel 160+8 t=1 type.*map.bucket[string]int+0
type.*map.bucket[string]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 c9 be cc 9c 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*map.bucket[string]int-+0
        rel 48+8 t=1 type.noalg.map.bucket[string]int+0
type..namedata.*map.hdr[string]int- SRODATA dupok size=22
        0x0000 00 00 13 2a 6d 61 70 2e 68 64 72 5b 73 74 72 69  ...*map.hdr[stri
        0x0010 6e 67 5d 69 6e 74                                ng]int
type.*map.hdr[string]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 9b ec 34 c9 08 08 08 36 00 00 00 00 00 00 00 00  ..4....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*map.hdr[string]int-+0
        rel 48+8 t=1 type.noalg.map.hdr[string]int+0
runtime.gcbits.2c SRODATA dupok size=1
        0x0000 2c                                               ,
type..namedata.count- SRODATA dupok size=8
        0x0000 00 00 05 63 6f 75 6e 74                          ...count
type..namedata.flags- SRODATA dupok size=8
        0x0000 00 00 05 66 6c 61 67 73                          ...flags
type..namedata.B. SRODATA dupok size=4
        0x0000 01 00 01 42                                      ...B
type..namedata.noverflow- SRODATA dupok size=12
        0x0000 00 00 09 6e 6f 76 65 72 66 6c 6f 77              ...noverflow
type..namedata.hash0- SRODATA dupok size=8
        0x0000 00 00 05 68 61 73 68 30                          ...hash0
type..namedata.buckets- SRODATA dupok size=10
        0x0000 00 00 07 62 75 63 6b 65 74 73                    ...buckets
type..namedata.oldbuckets- SRODATA dupok size=13
        0x0000 00 00 0a 6f 6c 64 62 75 63 6b 65 74 73           ...oldbuckets
type..namedata.nevacuate- SRODATA dupok size=12
        0x0000 00 00 09 6e 65 76 61 63 75 61 74 65              ...nevacuate
type..namedata.extra- SRODATA dupok size=8
        0x0000 00 00 05 65 78 74 72 61                          ...extra
type.noalg.map.hdr[string]int SRODATA dupok size=296
        0x0000 30 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  0.......0.......
        0x0010 35 46 99 12 02 08 08 19 00 00 00 00 00 00 00 00  5F..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 09 00 00 00 00 00 00 00 09 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0090 12 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x00a0 00 00 00 00 00 00 00 00 14 00 00 00 00 00 00 00  ................
        0x00b0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x00c0 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x00d0 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
        0x00e0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x00f0 30 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  0...............
        0x0100 00 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  ........@.......
        0x0110 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0120 50 00 00 00 00 00 00 00                          P.......
        rel 32+8 t=1 runtime.gcbits.2c+0
        rel 40+4 t=5 type..namedata.*map.hdr[string]int-+0
        rel 44+4 t=6 type.*map.hdr[string]int+0
        rel 48+8 t=1 type..importpath..+0
        rel 56+8 t=1 type.noalg.map.hdr[string]int+80
        rel 80+8 t=1 type..namedata.count-+0
        rel 88+8 t=1 type.int+0
        rel 104+8 t=1 type..namedata.flags-+0
        rel 112+8 t=1 type.uint8+0
        rel 128+8 t=1 type..namedata.B.+0
        rel 136+8 t=1 type.uint8+0
        rel 152+8 t=1 type..namedata.noverflow-+0
        rel 160+8 t=1 type.uint16+0
        rel 176+8 t=1 type..namedata.hash0-+0
        rel 184+8 t=1 type.uint32+0
        rel 200+8 t=1 type..namedata.buckets-+0
        rel 208+8 t=1 type.*map.bucket[string]int+0
        rel 224+8 t=1 type..namedata.oldbuckets-+0
        rel 232+8 t=1 type.*map.bucket[string]int+0
        rel 248+8 t=1 type..namedata.nevacuate-+0
        rel 256+8 t=1 type.uintptr+0
        rel 272+8 t=1 type..namedata.extra-+0
        rel 280+8 t=1 type.unsafe.Pointer+0
runtime.strhash·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.strhash+0
type..namedata.*map[string]int- SRODATA dupok size=18
        0x0000 00 00 0f 2a 6d 61 70 5b 73 74 72 69 6e 67 5d 69  ...*map[string]i
        0x0010 6e 74                                            nt
type.*map[string]int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 8e 5e 06 04 08 08 08 36 00 00 00 00 00 00 00 00  .^.....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*map[string]int-+0
        rel 48+8 t=1 type.map[string]int+0
type.map[string]int SRODATA dupok size=88
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 e5 db c8 4a 02 08 08 35 00 00 00 00 00 00 00 00  ...J...5........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0050 10 08 d0 00 0c 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*map[string]int-+0
        rel 44+4 t=6 type.*map[string]int+0
        rel 48+8 t=1 type.string+0
        rel 56+8 t=1 type.int+0
        rel 64+8 t=1 type.noalg.map.bucket[string]int+0
        rel 72+8 t=1 runtime.strhash·f+0
gclocals·263043c8f03e3241528dfae4e2812ef4 SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
gclocals·6d8d66cc8b183c896811c0201ffd5671 SRODATA dupok size=16
        0x0000 02 00 00 00 20 00 00 00 00 00 00 00 2c 00 00 00  .... .......,...
"".M06.stkobj SRODATA static size=40
        0x0000 02 00 00 00 00 00 00 00 00 ff ff ff ff ff ff ff  ................
        0x0010 00 00 00 00 00 00 00 00 30 ff ff ff ff ff ff ff  ........0.......
        0x0020 00 00 00 00 00 00 00 00                          ........
        rel 16+8 t=1 type.noalg.map.hdr[string]int+0
        rel 32+8 t=1 type.noalg.map.bucket[string]int+0
```


## go tool compile -S ssa_07.go
``` s
"".M07 STEXT size=159 args=0x18 locals=0x30 funcid=0x0
        0x0000 00000 (ssa_07.go:3)      TEXT    "".M07(SB), ABIInternal, $48-24
        0x0000 00000 (ssa_07.go:3)      MOVQ    (TLS), CX
        0x0009 00009 (ssa_07.go:3)      CMPQ    SP, 16(CX)
        0x000d 00013 (ssa_07.go:3)      PCDATA  $0, $-2
        0x000d 00013 (ssa_07.go:3)      JLS     149
        0x0013 00019 (ssa_07.go:3)      PCDATA  $0, $-1
        0x0013 00019 (ssa_07.go:3)      SUBQ    $48, SP
        0x0017 00023 (ssa_07.go:3)      MOVQ    BP, 40(SP)
        0x001c 00028 (ssa_07.go:3)      LEAQ    40(SP), BP
        0x0021 00033 (ssa_07.go:3)      FUNCDATA        $0, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
        0x0021 00033 (ssa_07.go:3)      FUNCDATA        $1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0021 00033 (ssa_07.go:4)      LEAQ    type.chan int(SB), AX
        0x0028 00040 (ssa_07.go:4)      MOVQ    AX, (SP)
        0x002c 00044 (ssa_07.go:4)      MOVQ    $1, 8(SP)
        0x0035 00053 (ssa_07.go:4)      PCDATA  $1, $0
        0x0035 00053 (ssa_07.go:4)      CALL    runtime.makechan(SB)
        0x003a 00058 (ssa_07.go:4)      MOVQ    16(SP), AX
        0x003f 00063 (ssa_07.go:4)      MOVQ    AX, "".m+32(SP)
        0x0044 00068 (ssa_07.go:5)      MOVQ    AX, (SP)
        0x0048 00072 (ssa_07.go:5)      LEAQ    ""..stmp_0(SB), CX
        0x004f 00079 (ssa_07.go:5)      MOVQ    CX, 8(SP)
        0x0054 00084 (ssa_07.go:5)      PCDATA  $1, $1
        0x0054 00084 (ssa_07.go:5)      CALL    runtime.chansend1(SB)
        0x0059 00089 (ssa_07.go:6)      MOVQ    $0, ""..autotmp_2+24(SP)
        0x0062 00098 (ssa_07.go:6)      MOVQ    "".m+32(SP), AX
        0x0067 00103 (ssa_07.go:6)      MOVQ    AX, (SP)
        0x006b 00107 (ssa_07.go:6)      LEAQ    ""..autotmp_2+24(SP), AX
        0x0070 00112 (ssa_07.go:6)      MOVQ    AX, 8(SP)
        0x0075 00117 (ssa_07.go:6)      PCDATA  $1, $0
        0x0075 00117 (ssa_07.go:6)      CALL    runtime.chanrecv1(SB)
        0x007a 00122 (ssa_07.go:7)      MOVQ    $0, "".~r0+56(SP)
        0x0083 00131 (ssa_07.go:7)      XORPS   X0, X0
        0x0086 00134 (ssa_07.go:7)      MOVUPS  X0, "".~r0+64(SP)
        0x008b 00139 (ssa_07.go:7)      MOVQ    40(SP), BP
        0x0090 00144 (ssa_07.go:7)      ADDQ    $48, SP
        0x0094 00148 (ssa_07.go:7)      RET
        0x0095 00149 (ssa_07.go:7)      NOP
        0x0095 00149 (ssa_07.go:3)      PCDATA  $1, $-1
        0x0095 00149 (ssa_07.go:3)      PCDATA  $0, $-2
        0x0095 00149 (ssa_07.go:3)      CALL    runtime.morestack_noctxt(SB)
        0x009a 00154 (ssa_07.go:3)      PCDATA  $0, $-1
        0x009a 00154 (ssa_07.go:3)      JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 82  eH..%....H;a....
        0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
        0x0020 28 48 8d 05 00 00 00 00 48 89 04 24 48 c7 44 24  (H......H..$H.D$
        0x0030 08 01 00 00 00 e8 00 00 00 00 48 8b 44 24 10 48  ..........H.D$.H
        0x0040 89 44 24 20 48 89 04 24 48 8d 0d 00 00 00 00 48  .D$ H..$H......H
        0x0050 89 4c 24 08 e8 00 00 00 00 48 c7 44 24 18 00 00  .L$......H.D$...
        0x0060 00 00 48 8b 44 24 20 48 89 04 24 48 8d 44 24 18  ..H.D$ H..$H.D$.
        0x0070 48 89 44 24 08 e8 00 00 00 00 48 c7 44 24 38 00  H.D$......H.D$8.
        0x0080 00 00 00 0f 57 c0 0f 11 44 24 40 48 8b 6c 24 28  ....W...D$@H.l$(
        0x0090 48 83 c4 30 c3 e8 00 00 00 00 e9 61 ff ff ff     H..0.......a...
        rel 5+4 t=17 TLS+0
        rel 36+4 t=16 type.chan int+0
        rel 54+4 t=8 runtime.makechan+0
        rel 75+4 t=16 ""..stmp_0+0
        rel 85+4 t=8 runtime.chansend1+0
        rel 118+4 t=8 runtime.chanrecv1+0
        rel 150+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 64 69 72 61 30 39                                dira09
""..stmp_0 SRODATA static size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*chan int- SRODATA dupok size=12
        0x0000 00 00 09 2a 63 68 61 6e 20 69 6e 74              ...*chan int
type.*chan int SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 ed 7b ed 3b 08 08 08 36 00 00 00 00 00 00 00 00  .{.;...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*chan int-+0
        rel 48+8 t=1 type.chan int+0
type.chan int SRODATA dupok size=64
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 91 55 cb 71 0a 08 08 32 00 00 00 00 00 00 00 00  .U.q...2........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*chan int-+0
        rel 44+4 t=6 type.*chan int+0
        rel 48+8 t=1 type.int+0
gclocals·263043c8f03e3241528dfae4e2812ef4 SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
```
