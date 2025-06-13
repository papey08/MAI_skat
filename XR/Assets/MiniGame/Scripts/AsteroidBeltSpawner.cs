using UnityEngine;

public class AsteroidBeltSpawner : MonoBehaviour
{
    public GameObject asteroidPrefab;
    public Transform centerObject;

    public float innerRadius;
    public float outerRadius;
    public int asteroidCount;
    public float minScale;
    public float maxScale;

    private void Start()
    {
        SpawnAsteroidBelt();
    }

    private void SpawnAsteroidBelt()
    {
        for (int i = 0; i < asteroidCount; i++)
        {
            float angle = Random.Range(0f, Mathf.PI * 2f);
            float radius = Random.Range(innerRadius, outerRadius);
            Vector2 spawnPos = (Vector2)centerObject.position + new Vector2(Mathf.Cos(angle), Mathf.Sin(angle)) * radius;
            GameObject asteroid = Instantiate(asteroidPrefab, spawnPos, Quaternion.identity, transform);
            float scale = Random.Range(minScale, maxScale);
            asteroid.transform.localScale = new Vector3(scale, scale, 1f);
            asteroid.transform.rotation = Quaternion.Euler(0f, 0f, Random.Range(0f, 360f));
        }
    }
}
