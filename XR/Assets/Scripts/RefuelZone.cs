using UnityEngine;

public class RefuelZone : MonoBehaviour
{
    public float refuelRange;

    public bool IsPlayerInRange(Transform player)
    {
        float distance = Vector3.Distance(transform.position, player.position);
        return distance <= refuelRange;
    }

    void OnDrawGizmosSelected()
    {
        Gizmos.color = Color.green;
        Gizmos.DrawWireSphere(transform.position, refuelRange);
    }
}
