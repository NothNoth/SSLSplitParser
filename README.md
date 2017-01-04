# spparser

Parses an SSLSplit log file into descriptors and data chunks.

# spexplode

Explodes aloaded SSSplit structure to folders and files

# spextract

Basic main function parsing an SSLSplit log and exploding it into a given output folder

Building:

    cd SSLSplitParser/spextract/
    go get
    go build
  
Usage:

    ./spextract output_1483372352.log output
    219 SSLSplit chunks loaded
    # 0 from 172.16.42.185 to 212.27.48.2 ( 18 bytes )
    # 1 from 172.16.42.185 to 212.27.48.2 ( 18 bytes )
    # 2 from 172.16.42.185 to 64.233.166.109 ( 69 bytes )
    # 3 from 172.16.42.185 to 64.233.166.109 ( 68 bytes )
    # 4 from 172.16.42.185 to 64.233.166.109 ( 68 bytes )
    # 5 from 172.16.42.185 to 64.233.166.109 ( 68 bytes )
    # 6 from 172.16.42.185 to 64.233.166.109 ( 68 bytes )
    # 7 from 172.16.42.185 to 64.233.166.109 ( 68 bytes )
    # 8 from 172.16.42.185 to 40.101.60.226 ( 160 bytes )
    # 9 from 172.16.42.185 to 40.101.60.226 ( 160 bytes )
    # 10 from 2.16.117.16 to 172.16.42.112 ( 182 bytes )
    # 11 from 172.16.42.112 to 2.16.117.16 ( 1448 bytes )
    # 12 from 172.16.42.112 to 2.16.117.16 ( 1448 bytes )
    ...
    
    ls -R output/
    output/:
    172.16.42.112-104.85.54.199    185.33.222.249-172.16.42.112  216.52.1.12-172.16.42.112     40.101.60.226-172.16.42.185   87.250.251.119-172.16.42.112
    172.16.42.112-104.94.193.130   189.1.173.100-172.16.42.112   216.58.198.194-172.16.42.112  52.19.220.125-172.16.42.112   88.212.243.76-172.16.42.112
    172.16.42.112-104.96.17.199    204.79.197.200-172.16.42.112  216.58.204.104-172.16.42.112  52.28.158.247-172.16.42.112   88.221.83.154-172.16.42.112
    172.16.42.112-109.248.237.34   207.46.194.14-172.16.42.112   216.58.204.110-172.16.42.112  62.152.54.250-172.16.42.112   89.249.22.206-172.16.42.112
    172.16.42.112-130.211.83.208   212.224.118.36-172.16.42.112  23.205.82.248-172.16.42.112   63.215.202.72-172.16.42.112   89.249.22.211-172.16.42.112
    172.16.42.112-137.116.198.146  212.224.124.70-172.16.42.112  23.97.187.123-172.16.42.112   64.233.166.109-172.16.42.185  89.249.22.218-172.16.42.112
    172.16.42.112-138.201.230.88   212.27.48.2-172.16.42.185     31.13.92.14-172.16.42.112     77.109.85.18-172.16.42.112    91.228.155.61-172.16.42.112
    178.250.0.80-172.16.42.112     213.19.162.79-172.16.42.112   37.140.238.33-172.16.42.112   81.222.128.15-172.16.42.112   95.131.137.7-172.16.42.112
    178.250.2.74-172.16.42.112     2.16.117.16-172.16.42.112     37.157.4.15-172.16.42.112     85.17.189.108-172.16.42.112

    output/172.16.42.112-104.85.54.199:
    1483372410.raw  1483372410.txt  1483372415.raw  1483372415.txt  1483372416.raw  1483372416.txt

    output/172.16.42.112-104.94.193.130:
    1483372410.raw  1483372410.txt

    output/172.16.42.112-104.96.17.199:
    1483372416.raw  1483372416.txt

    output/172.16.42.112-109.248.237.34:
    1483372416.raw  1483372416.txt

    output/172.16.42.112-130.211.83.208:
    1483372419.raw  1483372419.txt

    output/172.16.42.112-137.116.198.146:
    1483372418.raw  1483372418.txt  1483372419.raw  1483372419.txt  1483372420.raw  1483372420.txt

    output/172.16.42.112-138.201.230.88:
    1483372415.raw  1483372415.txt  1483372416.raw  1483372416.txt

    output/178.250.0.80-172.16.42.112:
    1483372414.raw  1483372414.txt

    output/178.250.2.74-172.16.42.112:
    1483372410.raw  1483372410.txt

    output/185.33.222.249-172.16.42.112:
    1483372419.raw  1483372419.txt

    output/189.1.173.100-172.16.42.112:
    1483372413.raw  1483372413.txt  1483372414.raw  1483372414.txt

    output/204.79.197.200-172.16.42.112:
    1483372414.raw  1483372414.txt

    output/207.46.194.14-172.16.42.112:
    1483372415.raw  1483372415.txt  1483372416.raw  1483372416.txt

    output/212.224.118.36-172.16.42.112:
    1483372414.raw  1483372414.txt

    output/212.224.124.70-172.16.42.112:
    1483372413.raw  1483372413.txt

    output/212.27.48.2-172.16.42.185:
    1483372384.raw  1483372384.txt

    output/213.19.162.79-172.16.42.112:
    1483372414.raw  1483372414.txt

    output/2.16.117.16-172.16.42.112:
    1483372397.raw  1483372397.txt

    output/216.52.1.12-172.16.42.112:
    1483372419.raw  1483372419.txt

    output/216.58.198.194-172.16.42.112:
    1483372414.raw  1483372414.txt  1483372419.raw  1483372419.txt

    output/216.58.204.104-172.16.42.112:
    1483372410.raw  1483372410.txt

    output/216.58.204.110-172.16.42.112:
    1483372414.raw  1483372414.txt

    output/23.205.82.248-172.16.42.112:
    1483372410.raw  1483372411.raw  1483372412.raw  1483372414.raw  1483372415.raw
    1483372410.txt  1483372411.txt  1483372412.txt  1483372414.txt  1483372415.txt

    output/23.97.187.123-172.16.42.112:
    1483372418.raw  1483372418.txt

    output/31.13.92.14-172.16.42.112:
    1483372413.raw  1483372413.txt

    output/37.140.238.33-172.16.42.112:
    1483372413.raw  1483372413.txt  1483372415.raw  1483372415.txt  1483372416.raw  1483372416.txt

    output/37.157.4.15-172.16.42.112:
    1483372416.raw  1483372416.txt

    output/40.101.60.226-172.16.42.185:
    1483372385.raw  1483372385.txt

    output/52.19.220.125-172.16.42.112:
    1483372414.raw  1483372414.txt

    output/52.28.158.247-172.16.42.112:
    1483372415.raw  1483372415.txt

    output/62.152.54.250-172.16.42.112:
    1483372416.raw  1483372416.txt

    output/63.215.202.72-172.16.42.112:
    1483372415.raw  1483372415.txt

    output/64.233.166.109-172.16.42.185:
    1483372384.raw  1483372384.txt  1483372385.raw  1483372385.txt

    output/77.109.85.18-172.16.42.112:
    1483372417.raw  1483372417.txt  1483372418.raw  1483372418.txt

    output/81.222.128.15-172.16.42.112:
    1483372416.raw  1483372416.txt

    output/85.17.189.108-172.16.42.112:
    1483372417.raw  1483372417.txt

    output/87.250.251.119-172.16.42.112:
    1483372413.raw  1483372413.txt

    output/88.212.243.76-172.16.42.112:
    1483372419.raw  1483372419.txt  1483372420.http  1483372420.raw  1483372420.txt

    output/88.221.83.154-172.16.42.112:
    1483372413.raw  1483372413.txt  1483372414.raw  1483372414.txt

    output/89.249.22.206-172.16.42.112:
    1483372417.raw  1483372417.txt

    output/89.249.22.211-172.16.42.112:
    1483372418.raw  1483372418.txt

    output/89.249.22.218-172.16.42.112:
    1483372417.raw  1483372417.txt

    output/91.228.155.61-172.16.42.112:
    1483372416.raw  1483372416.txt  1483372417.raw  1483372417.txt

    output/95.131.137.7-172.16.42.112:
    1483372416.raw  1483372416.txt