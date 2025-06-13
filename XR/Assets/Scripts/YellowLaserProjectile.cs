using UnityEngine;

public class YellowLaserProjectile : MonoBehaviour
{
    public float speed;
    public float lifetime;

    void Start()
    {
        Destroy(gameObject, lifetime);
    }

    void Update()
    {
        transform.Translate(Vector3.forward * speed * Time.deltaTime);
    }

    private void OnTriggerEnter(Collider other)
    {
        EnemyIdentifier enemy = other.GetComponent<EnemyIdentifier>();
        if (enemy != null && (enemy.enemyType == "BombardiroCrocodilo" || enemy.enemyType == "LaVaccaSaturnoSaturnita" || enemy.enemyType == "Glorbo"))
        {
            Destroy(other.gameObject);
        }
    }
}
