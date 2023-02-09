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

  
  надо скачать доп пакеты для работы с базой MySQL, 
  на терминале введем:
  
  go mod init ecommerceApp(только если не работает проект, или возникли с ним трудности, а так не надо)
  
      go get -u github.com/go-sql-driver/mysql 
  
  go get github.com/gorilla/sessions (только если они не скачаны, а так не надо, поидеи через мерч они будут у вас, поэтому сначала делайте мерч, даже если после мерча не будут эти пакеты, тока тогда скачиваем)
  
  go get github.com/jeypc/go-auth (только если они не скачаны, а так не надо, поидеи через мерч они будут у вас, поэтому сначала делайте мерч, даже если после мерча не будут эти пакеты, тока тогда скачиваем)
  
  go get golang.org/x/crypto/bcrypt (только если они не скачаны, а так не надо, поидеи через мерч они будут у вас, поэтому сначала делайте мерч, даже если после мерча не будут эти пакеты, тока тогда скачиваем)

    Не трогать ветку main. Работаем только со своими ветками
    
    Для слияние веток: git merge
    
    
   Для удобство с работой гит лучше скачать GoLand https://www.jetbrains.com/help/go/installation-guide.html
   . Для студентов дается лицензия на год. Для этого заходите в JetBrains и регаетесь как студент(SDU email). Затем у вас будет доступ к продукциям JetBrains.
