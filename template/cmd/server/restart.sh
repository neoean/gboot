if [ ! -f "./server" ];then
   echo "新文件不存在"
   else
   echo "文件存在"

   # shellcheck disable=SC2006
   PROCESS=`ps -ef | grep server_run |awk '{print $2}'`
   for i in $PROCESS
   do
     echo "Kill the $1 process [ $i ]"
      kill -9 $i
   done
   rm -rf server_run
   mv server server_run
fi

chmod 777 server_run

ntpdate time.nist.gov
export env=prod

cd ../..

./cmd/server/server_run >server_run.log
