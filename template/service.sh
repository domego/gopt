#!/bin/sh

ARGS=""
NAME=`cat NAME`
PID=`cat log/app.pid 2>/dev/null`
ACTION=$1

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
mkdir -p log

OS=`uname`
ECHO_ARGS="-e"
if [[ "$OS" = "Darwin" ]]; then
	ECHO_ARGS=""
fi
CRONTAB_DIR="$2"

function usage() {
	echo "usage: $0 [start|restart|reload|stop|run|status|version|check|install-check|uninstall-check]"
}

if [[ $# < 1 ]]; then
	usage
	exit 1
fi

function god() {
	GOD_NAME="bin/daemonize"
	if [[ -f $GOD_NAME ]]; then
		echo $GOD_NAME
		return
	fi
	case `uname` in
		Linux)
			GOD_NAME="bin/daemonize-linux"
			;;
		Darwin)
			GOD_NAME="bin/daemonize-darwin"
			;;
	esac
	echo $GOD_NAME
}

GOD=`god`

function stop() {
	kill $PID 2>/dev/null
	rm -rf log/app.pid
}

function start() {
	ps -p $PID >/dev/null 2>&1
	STOPPED=$?
	if [[ $STOPPED = 0 ]]; then
		echo $ECHO_ARGS "start [${GREEN}OK${NC}], service is running"
		return
	fi
	COUNT=`ulimit -n`
	LIMIT="10000"
	if [[ $COUNT -lt $LIMIT  ]]; then
		ulimit -n $LIMIT
	fi
	$GOD -a -e log/$NAME.err.log -o log/$NAME.log -p log/app.pid -c `pwd` `pwd`/bin/$NAME >/dev/null 2>&1
	sleep 0.2
	PID=`cat log/app.pid 2>/dev/null`
	STOPPED=$?
	if [[ "$STOPPED" = "0" ]]; then
		echo $ECHO_ARGS "start [${GREEN}OK${NC}]"
	else
		echo $ECHO_ARGS "start [${RED}FAIL${NC}]"
		exit 1
	fi
}

function restart() {
	ps -p $PID >/dev/null 2>&1
	STOPPED=$?
	if [[ "$STOPPED" = "0" ]]; then
		kill -HUP $PID
		echo $ECHO_ARGS "restart [${GREEN}OK${NC}]"
	else
		start
	fi
}

function check() {
	ps -p $PID >/dev/null 2>&1
	STOPPED=$?
	if [[ $STOPPED = 1 ]]; then
		echo "`date "+%Y-%m-%d %H:%M:%S"` restart $NAME"
		start
	fi
}

function install_check() {
	TEMP=`mktemp /tmp/crontab.XXXX`
	crontab -l > $TEMP
	KEY="$CRONTAB_DIR; ./service.sh check"
	grep "$KEY" $TEMP >/dev/null 2>&1
	NOT_INSTALL=$?
	if [[ $NOT_INSTALL = 1 ]]; then
		echo "*/1 * * * * cd $CRONTAB_DIR; ./service.sh check >> log/check.log" >> $TEMP
		crontab $TEMP
	fi
	rm $TEMP
}

function uninstall_check() {
	TEMP=`mktemp /tmp/crontab.XXXX`
	KEY="`pwd`; ./service.sh check"
	crontab -l | grep -v "$KEY"  > $TEMP
	crontab $TEMP
	rm $TEMP
}

case $1 in
	run)
		stop
		bin/$NAME
		;;
	start)
		start
	;;
	restart)
		restart
		;;
	reload)
		kill -s USR1 $PID
		;;
	stop)
		stop
		;;
	status)
		ps -fp $PID 2>/dev/null
		;;
	check)
		check
		;;
	install-check)
		install_check
		;;
	uninstall-check)
		uninstall_check
		;;
	version)
		version
		;;
	dev)
		stop
		make build
		if [[ $? = 0 ]]; then
      bin/$NAME
		fi
		;;
	*)
		usage
		exit 1
		;;
esac
