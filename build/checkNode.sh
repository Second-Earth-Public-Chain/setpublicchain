# Script usage：
# ./checkNode.sh 1 10 : Detect the latest block synchronization height of miner 1-miner 10

oldifs="$IFS"
IFS=$'\n'

function checkOneNode ()
{
    number=0
    logFile="logs/minernodetest"$1".log"
    for line in `tail $logFile`
    do
            len=${#line}
            if [[ $len -lt 100 ]]; then
                    continue
            fi
            line=`expr substr $line 50 200`
            tmpNumber=`expr $line : '.*number=\(.*\) '`
            if [[ $tmpNumber -gt $number ]]; then
                    number=$tmpNumber
            fi
    done
    echo minernodetest$1:$number
}

if [[ $# -ne 2 ]]; then
        echo "command=>$0, no parameters"
        exit 1
fi

startNodeNum=$1
while(( $startNodeNum<=$2 ))
do
        checkOneNode $startNodeNum
        let "startNodeNum++"
done


IFS="$oldifs"
