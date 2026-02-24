package utility

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Base58Encode 将字节数组编码为 base58 字符串
func Base58Encode(input []byte) string {
	if len(input) == 0 {
		return ""
	}
	x := new(big.Int).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)
	var result []byte
	for x.Cmp(zero) > 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}
	// 前导零字节 -> '1'
	for _, b := range input {
		if b != 0 {
			break
		}
		result = append(result, base58Alphabet[0])
	}
	// 反转
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

// Base58Decode 将 base58 字符串解码为字节数组
func Base58Decode(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("空的输入")
	}
	result := big.NewInt(0)
	base := big.NewInt(58)
	for _, c := range input {
		idx := strings.IndexRune(base58Alphabet, c)
		if idx < 0 {
			return nil, fmt.Errorf("包含无效字符: %c", c)
		}
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(idx)))
	}
	decoded := result.Bytes()
	// 恢复前导零
	for _, c := range input {
		if c != rune(base58Alphabet[0]) {
			break
		}
		decoded = append([]byte{0}, decoded...)
	}
	return decoded, nil
}

// EncodeWallets 将钱包地址列表编码为分享码
func EncodeWallets(wallets []string) (string, error) {
	if len(wallets) == 0 {
		return "", fmt.Errorf("没有可导出的钱包")
	}
	// 每个地址 20 字节，前置 1 字节表示地址数量
	buf := make([]byte, 0, 1+len(wallets)*20)
	buf = append(buf, byte(len(wallets)))
	for _, w := range wallets {
		addr := strings.TrimPrefix(strings.ToLower(w), "0x")
		if len(addr) != 40 {
			return "", fmt.Errorf("无效的钱包地址: %s", w)
		}
		b, err := hex.DecodeString(addr)
		if err != nil {
			return "", fmt.Errorf("无效的钱包地址: %s", w)
		}
		buf = append(buf, b...)
	}
	return Base58Encode(buf), nil
}

// DecodeWallets 将分享码解码为钱包地址列表
func DecodeWallets(code string) ([]string, error) {
	data, err := Base58Decode(code)
	if err != nil {
		return nil, err
	}
	if len(data) < 1 {
		return nil, fmt.Errorf("无效的分享码")
	}
	count := int(data[0])
	body := data[1:]
	if len(body) != count*20 {
		return nil, fmt.Errorf("分享码数据长度不匹配，期望 %d 个地址", count)
	}
	wallets := make([]string, 0, count)
	for i := 0; i < count; i++ {
		addr := "0x" + hex.EncodeToString(body[i*20:(i+1)*20])
		wallets = append(wallets, addr)
	}
	return wallets, nil
}
