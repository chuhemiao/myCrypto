package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)
/*   对称加密  start  */
// des 加密实现   cbc 分组模式
// des 长度8字节

// 填充最后一个分组函数    对称加密 推荐使用aes

// 3des 每个字节8字节，一共24个字节   依然按8字节分配加密   三组数据不同加密更安全

// aes
// 加密要求长度 16、24、32字节  go中指定16字节长度
// 分组长度  16,32,24 三种字节   按密钥长度对应分配  只能16个字节

// 尾部填充  如果刚好是整数倍 直接增加一个分组 然后解密的时候去掉
// 如果不是整数倍 ，则补上字节，缺几位补几位


/*
	src -- 原始数据
	blockSize -- 每个分组的长度
 */

func paddingText(src []byte,blockSize int) []byte{
	// 最后一个分组要填充的多少字节
	padding := blockSize - len(src)%blockSize
	// 创建一个新的切片    每个字节的值为什么

	padText := bytes.Repeat([]byte{byte(padding)},padding)
	// 参数二为切片
	newText := append(src,padText...)

	// 返回新的字符串
	return newText

}

// 删除末尾填充的字节

func unPaddingText(src []byte) []byte{
	// 求出字符串的长度
	len := len(src)
	// 去除最后一个字符
	number := int(src[len-1])
	// 删除末尾字节
	newText := src[:len-number]
	return  newText
}

// des加密

func encryptDES(src,ikey []byte ) []byte {

	block , err := des.NewCipher(ikey)
	if err != nil{
		panic("err")
	}
	src = paddingText(src,block.BlockSize())
	// 密码分组模式
	iv := []byte("aaaabbbb")
	blockNode := cipher.NewCBCEncrypter(block,iv)
	dst := make([]byte,len(src))
	blockNode.CryptBlocks(dst,src)
	return dst
}

// des 解密

func decryptDES(src,ikey []byte ) []byte {
	block , err := des.NewCipher(ikey)
	if err != nil{
		panic("err")
	}
	iv := []byte("aaaabbbb")
	blockMode := cipher.NewCBCDecrypter(block,iv)
	// 解密
	blockMode.CryptBlocks(src,src)
	// 去填充
	newText := unPaddingText(src)

	return newText
}

// 3des 加密操作

func encrypt3DES(src,key [] byte ) [] byte{
	// 创建一个使用3des加密得算法
	block , err := des.NewTripleDESCipher(key)
	if err != nil{
		panic(err)
	}
	// 对最后一个明文数据填充
	src = paddingText(src,block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block,key[:block.BlockSize()])
	// 加密连续得数据
	blockMode.CryptBlocks(src,src)
	return  src
}

// 3des解密操作

func decrypt3DES(src,key[]byte ) []byte {
	block, err := des.NewTripleDESCipher(key)
	if err !=nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block,key[:block.BlockSize()])
	blockMode.CryptBlocks(src,src)
	// 去掉最后一组得填充数据
	unPaddingText(src)
	return src
}

// aes 加密  只有16字节得密钥

func encryptAES(src,key[]byte) []byte{
	// 创建一个aes算法的ciper.block
	block,err  := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 填充数据
	src = paddingText(src,block.BlockSize())
	// Aes加密得blockmode接口
	blockMode := cipher.NewCBCEncrypter(block,key)
	// 加密数据
	blockMode.CryptBlocks(src,src)
	return  src
}

// aes 解密
func decryptAES(src,key[]byte) []byte{
	block ,err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block,key)
	blockMode.CryptBlocks(src,src)
	src = unPaddingText(src)
	return src
}

/*   对称加密 end   */

