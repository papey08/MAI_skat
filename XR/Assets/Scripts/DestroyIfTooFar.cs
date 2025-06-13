using UnityEngine;

public class DestroyIfTooFar : MonoBehaviour
{
    public float maxDistance;
    public GameObject player;

    void Update()
    {
        if (player == null) return;

        float distance = Vector3.Distance(transform.position, player.transform.position);

        if (distance > maxDistance)
        {
            Destroy(player);
        }
    }
}
