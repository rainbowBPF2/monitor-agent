# monitor-agent
Monitor Agent for OPS team.

1. Monitor Agent is mainly used for running a command periodicity.
   
   a) support running command or shell like cron job.
   
   b) support data encrypted and decrypted with AES algorithm, when deal with shell file.
   
   c) support self-define agent configuration, using conf/app.conf.
   
   d) support fast-deploy without no dependency(No JVM, Python or other Runtime, just go build and ready to deploy).
   
   
2. Easy Start
 
   a) git clone it.
   
   b) config go-path include this project.
   
   c) Terminal environment, cd project directory, go build .
   
   d) Get target file agent or agent.exe and conf.
   
   e) Run it: ./agent or agent.exe 
   
   
3. Project Note:
   a) Just support configuration by config file now(using conf/app.conf).
    
   b) If you want decrypt an encrypted shell. You need use the same encrypt salt key to encrypt it firstly.
    
   c) config file location: src/conf/app.conf, which is well described with related comment.
    
    
    
          
