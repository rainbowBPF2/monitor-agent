# monitor-agent
Monitor Agent for OPS team.

1. Monitor Agent is mainly used for running a command periodicity.
   
   a) support running command or shell like cron job.
   
   b) support data encrypted and decrypted with AES algorithm, when deal with shell file.
   
   c) support self-define agent configuration, using conf/app.conf.
   
   d) support fast-deploy without no dependency(No JVM, Python or other Runtime, just go build and ready to deploy).
   
   
2. Easy Start
 
   1) git clone it.
   
   2) config go-path include this project.
   
   3) Terminal environment, cd project directory, go build .
   
   4) Get target file agent or agent.exe and conf.
   
   5) Run it: ./agent or agent.exe 
   
   
3. Project Note:
   1) Just support configuration by config file now(using conf/app.conf).
    
   2) If you want decrypt an encrypted shell. You need use the same encrypt salt key to encrypt it firstly.
    
   3) config file location: src/conf/app.conf, which is well described with related comment.
    
    
    
          
