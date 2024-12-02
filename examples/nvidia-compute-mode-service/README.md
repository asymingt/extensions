# NVidia Compute Mode Service

NVidia GPU compute mode is configured at runtime through the `nvidia-smi` tool, and cannot persist across reboots. See:
https://forums.developer.nvidia.com/t/nvidia-smi-how-to-make-compute-mode-permanent-compute-mode-reverts-to-0-after-reboot/18877

This is a Talos system extension that does this configuration as a service. It basically does this:

```sh
nvidia-smi -i 0 -c EXCLUSIVE_PROCESS
nvidia-cuda-mps-control -d 
```
