# ecommerceApp
  В проекте учавствуют:
    
      Rustem Nygmet 200103424
      Telzhan Mukhadas 200103387
      Yerlan Yeldesov 200103084
      Berik Serikbay 200103348
      Aina Akimniyazova 200103139
      
      
  для клона:
  
      git clone https://github.com/Krasav4ik01/ecommerceApp.git -b yourBranch
  
  для комита и пуша:

      git commit -m "just commit"

      git push
      
      
  Для слияние веток:
 
        git merge

  
  надо скачать доп пакеты для работы с базой MySQL, 
  на терминале введем:
  
  
  
      go get -u github.com/go-sql-driver/mysql 
  
  
  А эти команды введем только если они не скачаны, а так не надо, поидеи через мерч они будут у вас, поэтому сначала делайте мерч, даже если после мерча не будут эти пакеты, тока тогда скачиваем:
  
  
      go get -u github.com/go-sql-driver/mysql 
      
      go get golang.org/x/crypto/bcrypt
      
      go get github.com/gorilla/sessions 
  
      

 Не трогать ветку main. Работаем только со своими ветками
    
        
 
 Доп команды, если проект не запустился, сначала удалаем go.mod, затем пишем на терминале:
        
        go mod init github.com/jeypc/go-auth
        
        
    и затем введем команды выше /\
        
        
        
    
    
   Для удобство с работой гит лучше скачать GoLand https://www.jetbrains.com/help/go/installation-guide.html
   . Для студентов дается лицензия на год. Для этого заходите в JetBrains и регаетесь как студент(SDU email). Затем у вас будет доступ к продукциям JetBrains.
