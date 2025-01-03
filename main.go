package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func flattenDir(srcDir, dstDir string) error {
    // 检查源目录是否存在
    if _, err := os.Stat(srcDir); os.IsNotExist(err) {
        return fmt.Errorf("source directory does not exist")
    }

    // 创建目标目录（如果不存在）
    if err := os.MkdirAll(dstDir, 0755); err != nil {
        return fmt.Errorf("failed to create destination directory: %v", err)
    }

    // 遍历源目录
    err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // 跳过目录
        if info.IsDir() {
            return nil
        }

        fmt.Printf("Processing file: %s\n", path)

        // 构建目标文件路径
        fileName := filepath.Base(path)
        dstPath := filepath.Join(dstDir, fileName)

        // 复制文件
        srcFile, err := os.Open(path)
        if err != nil {
            return fmt.Errorf("failed to open source file: %v", err)
        }
        defer srcFile.Close()

        dstFile, err := os.Create(dstPath)
        if err != nil {
            return fmt.Errorf("failed to create destination file: %v", err)
        }
        defer dstFile.Close()

        _, err = io.Copy(dstFile, srcFile)
        if err != nil {
            return fmt.Errorf("failed to copy file: %v", err)
        }

        // 复制文件权限和时间戳
        srcInfo, err := os.Stat(path)
        if err != nil {
            return fmt.Errorf("failed to get source file info: %v", err)
        }

        err = os.Chmod(dstPath, srcInfo.Mode())
        if err != nil {
            return fmt.Errorf("failed to set file permissions: %v", err)
        }

        return nil
    })

    if err != nil {
        return fmt.Errorf("error walking through directory: %v", err)
    }

    return nil
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: flattendir src_dir dst_dir")
        os.Exit(1)
    }

    srcDir := os.Args[1]
    dstDir := os.Args[2]

    if err := flattenDir(srcDir, dstDir); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Flatten directory successfully.")
}