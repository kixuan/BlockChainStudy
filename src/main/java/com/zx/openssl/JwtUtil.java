package com.zx.openssl;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;

import java.security.KeyFactory;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.util.Base64;

/**
 * JWT工具类
 * author: Kixuan
 */
public class JwtUtil {

    private static final String PRIVATE_KEY = "-----BEGIN PRIVATE KEY-----...-----END PRIVATE KEY-----";
    private static final String PUBLIC_KEY = "-----BEGIN PUBLIC KEY-----...-----END PUBLIC KEY-----";

    // 从PEM格式的字符串中解析并返回私钥
    private static PrivateKey getPrivateKey() throws Exception {
        String privateKeyPEM = PRIVATE_KEY.replace("-----BEGIN PRIVATE KEY-----", "")
                .replace("-----END PRIVATE KEY-----", "")
                .replaceAll("\\s", "");
        byte[] encoded = Base64.getDecoder().decode(privateKeyPEM);
        PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(encoded);
        KeyFactory keyFactory = KeyFactory.getInstance("RSA");
        return keyFactory.generatePrivate(keySpec);
    }

    // 从PEM格式的字符串中解析并返回公钥
    private static PublicKey getPublicKey() throws Exception {
        String publicKeyPEM = PUBLIC_KEY.replace("-----BEGIN PUBLIC KEY-----", "")
                .replace("-----END PUBLIC KEY-----", "")
                .replaceAll("\\s", "");
        byte[] encoded = Base64.getDecoder().decode(publicKeyPEM);
        X509EncodedKeySpec keySpec = new X509EncodedKeySpec(encoded);
        KeyFactory keyFactory = KeyFactory.getInstance("RSA");
        return keyFactory.generatePublic(keySpec);
    }

    // 使用私钥生成JWT令牌
    public static String generateToken(String subject) throws Exception {
        return Jwts.builder()
                .setSubject(subject)
                .signWith(getPrivateKey(), SignatureAlgorithm.RS256)
                .compact();
    }

    // 使用公钥验证JWT令牌
    public static boolean validateToken(String token) throws Exception {
        Jwts.parser()
                .setSigningKey(getPublicKey())
                .build()
                .parseClaimsJws(token);
        return true;
    }
}