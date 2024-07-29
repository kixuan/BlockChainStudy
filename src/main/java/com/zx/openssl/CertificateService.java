package com.zx.openssl;

import org.springframework.stereotype.Component;

import java.io.BufferedReader;
import java.io.InputStreamReader;

@Component
public class CertificateService {

    // 生成证书
    public void generateCertificate(String certificateName) throws Exception {
        String command = "openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj \"/CN=" + certificateName + "\"";
        executeCommand(command);
    }

    private void executeCommand(String command) throws Exception {
        Process process = Runtime.getRuntime().exec(command);
        BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
        String line;
        while ((line = reader.readLine()) != null) {
            System.out.println(line);
        }
        process.waitFor();
    }

    // 私钥签名
    public void signData(String data, String privateKeyFile) throws Exception {
        String command = "echo \"" + data + "\" | openssl dgst -sha256 -sign " + privateKeyFile + " -out data.sig";
        executeCommand(command);
    }

    // 验签
    public void verifySignature(String data, String signatureFile, String publicKeyFile) throws Exception {
        String command = "echo \"" + data + "\" | openssl dgst -sha256 -verify " + publicKeyFile + " -signature " + signatureFile;
        executeCommand(command);
    }
}
