## Запуск

#!/usr/bin/env bash
mkdir -p rl_env
# сохранить env.py, policy.py, train_ppo.py в папку rl_env/
cd rl_env
python3 train_ppo.py --episodes 500
