using UnityEngine;
using System.Collections.Generic;

public class AbsorbableSpawner : MonoBehaviour
{
    public GameObject absorbablePrefab;
    public int spawnCount;

    public float spawnThresholdScaleX;
    public float destroyThresholdScaleX;

    public float spawnRange;

    private Transform blackHole;
    private List<GameObject> spawnedObjects = new List<GameObject>();
    private bool hasSpawned = false;

    private void Start()
    {
        blackHole = GameObject.FindGameObjectWithTag("Player").transform;
    }

    private void Update()
    {
        float currentScaleX = blackHole.localScale.x;

        if (!hasSpawned && currentScaleX >= spawnThresholdScaleX)
        {
            SpawnAbsorbables();
            hasSpawned = true;
        }

        if (hasSpawned && currentScaleX >= destroyThresholdScaleX)
        {
            DestroyAbsorbables();
            hasSpawned = false;
        }
    }

    private void SpawnAbsorbables()
    {
        for (int i = 0; i < spawnCount; i++)
        {
            Vector2 spawnPos = GetSpawnPosition();
            Quaternion rotation = Quaternion.Euler(0f, 0f, Random.Range(0f, 360f));
            GameObject obj = Instantiate(absorbablePrefab, spawnPos, rotation);
            spawnedObjects.Add(obj);
        }
    }

    private Vector2 GetSpawnPosition()
    {
        float angle = Random.Range(0f, Mathf.PI * 2f);
        float radius = Mathf.Sqrt(Random.Range(0f, 1f)) * spawnRange;
        return new Vector2(Mathf.Cos(angle), Mathf.Sin(angle)) * radius;
    }

    private void DestroyAbsorbables()
    {
        foreach (var obj in spawnedObjects)
        {
            if (obj != null)
                Destroy(obj);
        }
        spawnedObjects.Clear();
    }
}
