using UnityEngine;

public class AbsorbableMotion : MonoBehaviour
{
    private float moveSpeed = 0.5f;
    private float rotationSpeed = 2f;
    private float directionChangeInterval = 30f;

    private Vector2 moveDirection;
    private float directionTimer;

    private void Start()
    {
        PickNewDirection();
        directionTimer = directionChangeInterval;
    }

    private void Update()
    {
        transform.Translate(moveDirection * moveSpeed * Time.deltaTime, Space.World);
        transform.Rotate(Vector3.forward * rotationSpeed * Time.deltaTime);
        directionTimer -= Time.deltaTime;
        if (directionTimer <= 0f)
        {
            PickNewDirection();
            directionTimer = directionChangeInterval;
        }
    }

    private void PickNewDirection()
    {
        float angle = Random.Range(0f, Mathf.PI * 2f);
        moveDirection = new Vector2(Mathf.Cos(angle), Mathf.Sin(angle)).normalized;
        rotationSpeed = Random.Range(-40f, 40f);
    }
}
