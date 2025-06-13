using UnityEngine;

public class RedLaserProjectile : MonoBehaviour
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
        if (enemy != null && enemy.enemyType == "BombardiroCrocodilo")
        {
            TryAddBlueLaserCharge();
            Destroy(other.gameObject);
            Destroy(gameObject);
        }
    }

    void TryAddBlueLaserCharge()
    {
        GameObject player = GameObject.FindGameObjectWithTag("Player");
        if (player != null)
        {
            ShipController ship = player.GetComponent<ShipController>();
            if (ship != null)
            {
                ship.AddBlueLaserCharge();
            }
        }
    }
}
