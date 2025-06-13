using UnityEngine;

public class OrbitInPlane : MonoBehaviour
{
    public Transform target;
    public float orbitSpeed;

    private float initialZ;

    void Start()
    {
        if (target == null)
        {
            Debug.LogError("OrbitInPlane: Target not assigned.");
            enabled = false;
            return;
        }

        initialZ = transform.position.z;
    }

    void Update()
    {
        if (target == null) return;

        Vector3 currentPos = transform.position;
        currentPos.z = initialZ;

        transform.RotateAround(
            new Vector3(target.position.x, target.position.y, initialZ),
            Vector3.forward,
            orbitSpeed * Time.deltaTime
        );
        Vector3 pos = transform.position;
        pos.z = initialZ;
        transform.position = pos;
    }
}
