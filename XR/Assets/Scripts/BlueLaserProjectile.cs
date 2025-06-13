using UnityEngine;

public class BlueLaserProjectile : MonoBehaviour
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
        if (enemy != null && (enemy.enemyType == "BombardiroCrocodilo" || enemy.enemyType == "LaVaccaSaturnoSaturnita") )
        {
            TryAddYellowLaserCharge();
            Destroy(other.gameObject);
            Destroy(gameObject);
        }
    }

    void TryAddYellowLaserCharge()
    {
        GameObject player = GameObject.FindGameObjectWithTag("Player");
        if (player != null)
        {
            ShipController ship = player.GetComponent<ShipController>();
            if (ship != null)
            {
                ship.AddYellowLaserCharge();
            }
        }
    }
}
