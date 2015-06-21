package scriptcodes

func main() {

}

//Contants
const OP_0 = byte(0)
const OP_FALSE = byte(0)
const OP_PUSHDATA1 = byte(76)
const OP_PUSHDATA2 = byte(77)
const OP_PUSHDATA4 = byte(78)
const OP_1NEGATE = byte(79)
const OP_1 = byte(80)
const OP_TRUE = byte(80)
const OP_2 = byte(82)
const OP_3 = byte(83)
const OP_4 = byte(84)
const OP_5 = byte(85)
const OP_6 = byte(86)
const OP_7 = byte(87)
const OP_8 = byte(88)
const OP_9 = byte(89)
const OP_10 = byte(90)
const OP_11 = byte(91)
const OP_12 = byte(92)
const OP_13 = byte(93)
const OP_14 = byte(94)
const OP_15 = byte(95)
const OP_16 = byte(96)

//Flow Control
const OP_NOP = byte(97)
const OP_IF = byte(99)
const OP_NOTIF = byte(100)
const OP_ELSE = byte(103)
const OP_ENDIF = byte(104)
const OP_VERIFY = byte(105)
const OP_RETURN = byte(106)

//Stack
const OP_TOALTSTACK = byte(107)
const OP_FROMALTSTACK = byte(108)
const OP_IFDUP = byte(115)
const OP_DEPTH = byte(116)
const OP_DROP = byte(117)
const OP_DUP = byte(118)
const OP_NIP = byte(119)
const OP_OVER = byte(120)
const OP_PICK = byte(121)
const OP_ROLL = byte(122)
const OP_ROT = byte(123)
const OP_SWAP = byte(124)
const OP_TUCK = byte(125)
const OP_2DROP = byte(109)
const OP_2DUP = byte(110)
const OP_3DUP = byte(111)
const OP_2OVER = byte(112)
const OP_2ROT = byte(113)
const OP_2SWAP = byte(114)

//Splice
const OP_CAT = byte(126)    //disabled
const OP_SUBSTR = byte(127) //disabled
const OP_LEFT = byte(128)   //disabled
const OP_RIGHT = byte(129)  //disabled
const OP_SIZE = byte(130)

//Bitwise logic
const OP_INVERT = byte(131) //disabled
const OP_AND = byte(132)    //disabled
const OP_OR = byte(133)     //disabled
const OP_XOR = byte(134)    //disabled
const OP_EQUAL = byte(135)
const OP_EQUALVERIFY = byte(136)

//Arithmetic
const OP_1ADD = byte(139)
const OP_1SUB = byte(140)
const OP_2MUL = byte(141) //disabled
const OP_2DIV = byte(142) //disabled
const OP_NEGATE = byte(143)
const OP_ABS = byte(144)
const OP_NOT = byte(145)
const OP_0NOTEQUAL = byte(146)
const OP_ADD = byte(147)
const OP_SUB = byte(148)
const OP_MUL = byte(149)    //disabled
const OP_DIV = byte(150)    //disabled
const OP_MOD = byte(151)    //disabled
const OP_LSHIFT = byte(152) //disabled
const OP_RSHIFT = byte(153) //disabled
const OP_BOOLAND = byte(154)
const OP_BOOLOR = byte(155)
const OP_NUMEQUAL = byte(156)
const OP_NUMEQUALVERIFY = byte(157)
const OP_NUMNOTEQUAL = byte(158)
const OP_LESSTHAN = byte(159)
const OP_GREATERTHAN = byte(160)
const OP_LESSTHANOREQUAL = byte(161)
const OP_GREATERTHANOREQUAL = byte(162)
const OP_MIN = byte(163)
const OP_MAX = byte(164)
const OP_WITHIN = byte(165)

//Crypto
const OP_RIPEMD160 = byte(166)
const OP_SHA1 = byte(167)
const OP_SHA256 = byte(168)
const OP_HASH160 = byte(169)
const OP_HASH256 = byte(170)
const OP_CODESEPARATOR = byte(171)
const OP_CHECKSIG = byte(172)
const OP_CHECKSIGVERIFY = byte(173)
const OP_CHECKMULTISIG = byte(174)
const OP_CHECKMULTISIGVERIFY = byte(175)

//Pseudo-words
const OP_PUBKEYHASH = byte(253)
const OP_PUBKEY = byte(254)
const OP_INVALIDOPCODE = byte(255)

//Reserved words
const OP_RESERVED = byte(80)
const OP_VER = byte(98)
const OP_VERIF = byte(101)
const OP_VERNOTIF = byte(102)
const OP_RESERVED1 = byte(137)
const OP_RESERVED2 = byte(138)
const OP_NOP1 = byte(176)
const OP_NOP2 = byte(177)
const OP_NOP3 = byte(178)
const OP_NOP4 = byte(179)
const OP_NOP5 = byte(180)
const OP_NOP6 = byte(181)
const OP_NOP7 = byte(182)
const OP_NOP8 = byte(183)
const OP_NOP9 = byte(184)
const OP_NOP10 = byte(185)
