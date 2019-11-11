package main

import (
	"crypto/md5"
	"encoding/hex"
)

// 单向散列函数   哈希函数   =》  消息摘要函数
// 有一个输入和一个输出，输入一个消息后  通过函数计算出一个散列值，这个值固定长度，无论输入的值有多少字节 都是固定值  生成后的值可以当作一个文件的指纹，用来实现web服务器快速上传文件
// 数据具有一致性（完整性）
// 能够快速计算出散列值
// 具有单向性  不可逆计算

// 两个不同的消息产生同一个散列值的情况称为碰撞
// 单向散列使用场景

// 1.检测软件是否被篡改   下载后软件计算散列值  看是否被篡改
// 2.消息验证码  ssl/tsl  防止数据在传输过程中被篡改
// 3.数字签名
// 4.伪随机数生成器
// 5.一次性口令  无线网络一次性口令


// 常用的单向散列函数

// 1.MD4、MD5  返回长度16字节
// 2.sha-1、sha-224、sha-256、sha-384、sha-512
// 接受消息的长度有限制
// 散列值长度
// 		     比特数     字节数
//MD4        128bit     16bit
//MD5        128bit     16bit
//SHA-1      128bit     20bit
//SHA-224    224bit     28bit
//SHA-256    256bit     34bit
//SHA-384    384bit     48bit
//SHA-512    512bit     64bit

// 接受一个切片   返回16个字符  0-9 A-F
func ShaMd5Str(mdStr []byte ) string{
	// 计算数据的md5
	res := md5.Sum(mdStr)
	//myStr := fmt.Sprintf("%x",res)
	// 16进制直接转换  res为处理后的固定长度切片
	myStr := hex.EncodeToString(res[:])
	return myStr
}
// md5方式2   可以多次添加数据块 然后在计算hash值
func ShaMd5Str2(mdStr []byte ) string{
	// 得到一个hash接口
	res := md5.New()
	// 添加数据 第一种方式
	//io.WriteString(res,string(mdStr))
	// 添加数据的第二种方式
	res.Write(mdStr)
	myStr := res.Sum(nil)
	return hex.EncodeToString(myStr)
}

