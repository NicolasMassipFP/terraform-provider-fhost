# TODO: Fix me

```
(venv) gregory@goueg ~/greg/clone/terraform-provider-smc/examples/engines/single_fw/single_fw_switch_interfaces (develop) $ make plan
terraform plan
╷
│ Error: Unsupported block type
│ 
│   on main.tf line 38, in resource "smc_single_fw" "tf_single_fw":
│   38:     switch_interface {
│ 
│ Blocks of type "switch_interface" are not expected here.
╵
make: *** [Makefile:14: plan] Error 1
```