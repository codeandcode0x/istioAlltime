#!bin/bash
set -e;

sleep $INTERVAL_TIME;

if [ -d "$SQL_FILE" ]
then
  rm -rf $SQL_FILE;
fi

#e.coding.net/codingcorp/coding-dev/coding-dev-sql.git
if [ -z $GIT_BRANCH ]
then
  GIT_BRANCH="master";
fi
git clone --depth=1 -b $GIT_BRANCH  https://$GITUSER:$GITPASSWD@$GITURL;

# check
instance=`eval echo '$'"$MIGRATION"`;
mysql -h$instance -u$DBUSER -p$DBPASSWD --default-character-set=utf8mb4 -e "show databases;"

set +e;
function initDbJob(){
    for file in `ls $1` 
    do
        if [ -d $1"/"$file ]
        then
            initDbJob $1"/"$file;
        else
            if [[ "${file##*.}" == "sql" ]]
            then
              result=$(echo $1 | egrep "mariadb|mysql");
              if [[ "$result" != "" ]]
              then
                echo $1"/"$file;
                instance=`eval echo '$'"$MIGRATION"`;
                mysql -h$instance -u$DBUSER -p$DBPASSWD $2 --default-character-set=utf8mb4 < $1"/"$file; 
                sleep 1;
              fi
            fi
        fi
    done
}

#deploy service
function initSchemaJob(){
    for file in `ls $1` 
    do
        if [ -d $1"/"$file ]
        then
            initSchemaJob $1"/"$file;
        else
            if [[ "${file##*.}" == "sql" ]]
            then
              if [[ "${file%%.*}" == "schema" ]]
              then
                index=0$count;
                if [[ "$1" != "mariadb" ]]
                then
                  echo $1"/"$file;
                  instance=`eval echo '$'"$MIGRATION"`;
                  mysql -h$instance -u$DBUSER -p$DBPASSWD $2 --default-character-set=utf8mb4 < $1"/"$file; 
                  sleep 1;
                fi
              fi
            fi
        fi
    done
}

#deploy service
function initDataJob(){
    for file in `ls $1` 
    do
        if [ -d $1"/"$file ]
        then
            initDataJob $1"/"$file;
        else
            if [[ "${file##*.}" == "sql" ]]
            then
              if [ "${file%%.*}" != "schema" ] && [ "${file%%.*}" != "init_db" ]
              then
                echo $1"/"$file;
                instance=`eval echo '$'"$MIGRATION"`;
                mysql -h$instance -u$DBUSER -p$DBPASSWD $2 --default-character-set=utf8mb4 < $1"/"$file;
              fi
            fi
        fi
    done
}

#main 
function main() {
  #run jobs
  initDbJob $1 $2;
  initSchemaJob $1 $2;
  initDataJob $1 $2;
}

#run shell
main $1 $2;




