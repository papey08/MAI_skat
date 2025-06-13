using UnityEngine;

public class SatelliteOrbit : MonoBehaviour
{
    public Transform planet;
    public float orbitSpeed;

    private float orbitRadius;
    private float currentAngle;

    void Start()
    {
        if (planet == null)
        {
            enabled = false;
            return;
        }
        Vector2 direction = transform.position - planet.position;
        orbitRadius = direction.magnitude;
        currentAngle = Mathf.Atan2(direction.y, direction.x);
    }

    void Update()
    {
        if (planet == null) return;
        currentAngle += orbitSpeed * Mathf.Deg2Rad * Time.deltaTime;
        float x = planet.position.x + Mathf.Cos(currentAngle) * orbitRadius;
        float y = planet.position.y + Mathf.Sin(currentAngle) * orbitRadius;
        transform.position = new Vector3(x, y, transform.position.z);
    }
}
