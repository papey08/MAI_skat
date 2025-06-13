using UnityEngine;

public class EnemySpawner : MonoBehaviour
{
    public GameObject enemyPrefab;
    public float spawnInterval;
    public float spawnRadius;
    public bool isActive;
    public bool spawnOnce;

    private Transform player;
    private bool hasSpawned = false;
    private static bool bossKilled = false;

    void Start()
    {
        bossKilled = false;
        player = GameObject.FindGameObjectWithTag("Player")?.transform;
        InvokeRepeating(nameof(SpawnEnemy), 2f, spawnInterval);
    }

    void SpawnEnemy()
    {
        if (!isActive || bossKilled || enemyPrefab == null || player == null) return;
        if (spawnOnce && hasSpawned) return;

        Vector3 spawnPosition = transform.position + Random.onUnitSphere * spawnRadius;
        spawnPosition.y = transform.position.y;

        GameObject enemy = Instantiate(enemyPrefab, spawnPosition, Quaternion.Euler(0, 90, 0));
        hasSpawned = true;

        if (spawnOnce)
        {
            isActive = false;
            EnemyDeathNotifier notifier = enemy.AddComponent<EnemyDeathNotifier>();
            notifier.OnDeath += OnSpawnOnceEnemyKilled;
        }
    }

    void OnSpawnOnceEnemyKilled()
    {
        bossKilled = true;
    }

    public void RestartScene() {
        bossKilled = false;
        hasSpawned = false;
    }
}
