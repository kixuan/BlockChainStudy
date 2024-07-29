package com.zx.openssl;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.Resource;
import org.springframework.util.FileCopyUtils;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;

@RestController
public class CertificateController {

    @Value("classpath:certificates")
    private Resource certificatesDirectory;

    // 获取证书
    @GetMapping("/certificates/{filename}")
    public String getCertificate(@PathVariable String filename) throws IOException {
        Resource certificateFile = certificatesDirectory.createRelative(filename);
        byte[] fileData = FileCopyUtils.copyToByteArray(certificateFile.getInputStream());
        return new String(fileData);
    }

    public CertificateController(CertificateService certificateService) {
        this.certificateService = certificateService;
    }

    private final CertificateService certificateService;


    @PostMapping("/generate")
    public void generateCertificate(@RequestParam String certificateName) throws Exception {
        certificateService.generateCertificate(certificateName);
    }

    @PostMapping("/sign")
    public void signData(@RequestParam String data, @RequestParam String privateKeyFile) throws Exception {
        certificateService.signData(data, privateKeyFile);
    }

    @PostMapping("/verify")
    public void verifySignature(@RequestParam String data, @RequestParam String signatureFile, @RequestParam String publicKeyFile) throws Exception {
        certificateService.verifySignature(data, signatureFile, publicKeyFile);
    }
}
