# aliyunoss-stream-cli
阿里云廉价直播服务终端工具

### 环境变量
OSS_ENDPOINT
OSS_AK
OSS_SK
OSS_BUCKET

### 用法

                
    aliyunoss-stream-cli
        create
          -channel your-channel-name
          -playlist your-playlist-name

        info
          -channel your-channel-name

        list
          -prefix prefix- (optional)

        sign
          -channel your-channel-name
          -expiry 123 (seconds, default 10 years)

        del
          -channel your-channel-name
