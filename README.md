# u-protocol



# How to use

* After you make changes on `.proto` files, you need to generate `.pb.go` files.

* Make sure [Docker](https://docs.docker.com/docker-for-mac/install/) is up and running

* On the root of the repository, run:

  ```
  make
  ```

* It takes a while if this is first time you run.

* Sample output (first run):

  ```
  $ make
  docker build -t u-protocol .
  Sending build context to Docker daemon  3.381MB
  Step 1/9 : FROM golang:1.10
  1.10: Pulling from library/golang
  741437d97401: Pull complete 
  ...
  ...
  ...
  rm -rf go
  mkdir go
  for a in ./proto/utransferoutfis/transferoutfis.proto ./proto/utransferoutfis/service.proto ./proto/usmsgateway/service.proto ./proto/uzendeskgateway/service.proto ./proto/ucardfis/card_load.proto ./proto/ucardfis/card_transfer.proto ./proto/ucardfis/service.proto ./proto/ucardfis/balance.proto ./proto/ucardfis/card.proto ./proto/ukbankgateway/payment_confrm.proto ./proto/ukbankgateway/service.proto ./proto/uemailgateway/service.proto ./proto/ucszendesk/service.proto ./proto/utopupkbank/service.proto ./proto/utopupkbank/topup.proto ./proto/ui18n/service.proto ./proto/upushgateway/service.proto ./proto/ucardtransactionfis/service.proto ./proto/ucashoutkbank/cash_out_kbank.proto ./proto/uapp/service.proto ./proto/uapp/appuser.proto ./proto/ukyc/service.proto ./proto/umagiclink/service.proto ./proto/umagiclink/magiclink.proto ./proto/upayment/payment_method.proto ./proto/upayment/service.proto ./proto/common/money.proto ./proto/common/enum.proto ./proto/common/kbank.proto ./proto/common/event.proto ./proto/common/card.proto ./proto/common/user.proto ./proto/ukyckbank/service.proto ./proto/ucashout/cash_out.proto ./proto/ufile/file.proto ./proto/ureward/service.proto ./proto/ureward/reward.proto ./proto/uproduct/service.proto ./proto/uticket/service.proto ./proto/ucardcrifis/crifis.proto ./proto/ucard/card_load.proto ./proto/ucard/card_transfer.proto ./proto/ucard/service.proto ./proto/ucard/balance.proto ./proto/ucard/card.proto ./proto/ufx/fx.proto ./proto/ufx/service.proto ./proto/errors/service.proto ./proto/utransferout/service.proto ./proto/utransferout/transferout.proto ./proto/unotification/service.proto ./proto/uauth/service.proto ./proto/utopup/topup.proto ./proto/ucardtransaction/service.proto ./proto/uotp/service.proto ./proto/uotp/otp.proto ./proto/ufisgateway/te_adjust.proto ./proto/ufisgateway/fis/cardholderupdate.proto ./proto/ufisgateway/fis/cardapplication.proto ./proto/ufisgateway/fis/statuschange.proto ./proto/ufisgateway/fis/cardactivation.proto ./proto/ufisgateway/fis/cardload.proto ./proto/ufisgateway/fis/cardtransfer.proto ./proto/ufisgateway/fis/outgoing_messages.proto ./proto/ufisgateway/fis/balanceenquirymulticurrency.proto ./proto/ufisgateway/fis/cardunload.proto ./proto/ufisgateway/fis/managepin.proto ./proto/ufisgateway/fis/common.proto ./proto/ufisgateway/te_card_load_unload.proto ./proto/ufisgateway/outgoing_messages.proto ./proto/ufisgateway/te_auth_advice.proto ./proto/ufisgateway/service.proto ./proto/ufisgateway/te_fee.proto ./proto/ufisgateway/te_fin_advice.proto ./proto/ufisgateway/te_messages.proto ./proto/ufxkbank/service.proto ./proto/uuser/service.proto ./proto/uuser/user.proto ./proto/ucardcri/cri.proto ./proto/ucardcri/service.proto; do protoc -I=./proto --go_out=plugins=grpc:/go/src $a; done
  
  ```

# Note 
<b>When using Apple Silicon Macs</b>

Modify the first line of the Dockerfile to include --platform flag
```
FROM --platform=linux/amd64 golang:1.12
```

  