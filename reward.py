from rl_env.oracle_reward import OracleReferenceReward

# Инициализируем один раз перед обучением
oracle = OracleReferenceReward("oracle_prefix_table.json", node_capacity)

# Внутри rollout используем O(1) lookup
reward = oracle.get_reward(step_idx, current_stddev, agent_placed)
