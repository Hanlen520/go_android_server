import os
import argparse

parser = argparse.ArgumentParser()

parser.add_argument('-k', '--kill', action="store_true", help='kill server')
parser.add_argument('-b', '--build', action="store_true", help='build server')
parser.add_argument('-r', '--run', action="store_true", help='run server')


def kill_server():
    print('try to kill server ...')
    ps_output = os.popen('adb shell ps | grep go_android_server').read()
    if ps_output:
        print('old server detected.')
        pid = [i for i in ps_output.split(' ') if i][1]
        os.system('adb shell kill -9 ' + pid)
        print('kill finished.')
    else:
        print('no server detected.')

def build_server():
    print('GO to build server ...')
    os.system('CGO_ENABLED=0 GOARCH=arm GOOS=linux go build go_android_server.go')
    print('build finished.')

def push_server():
    print('push to device ...')
    os.system('adb push go_android_server /data/local/tmp')
    print('push finished.')

def run_server():
    print('try to start server ...')
    os.system('adb shell /data/local/tmp/go_android_server &')
    print('server started :)')


if __name__ == '__main__':
    outside_args = parser.parse_args()
    if outside_args.kill:
        kill_server()
    elif outside_args.build:
        build_server()
    else:
        kill_server()
        build_server()
        push_server()
        run_server()
