using UnityEngine;

public class BossDeathHandler : MonoBehaviour
{
    private void OnDestroy()
    {
        GameObject[] enemies = GameObject.FindGameObjectsWithTag("Enemy");
        foreach (GameObject enemy in enemies)
        {
            if (enemy != null && enemy != gameObject)
            {
                Destroy(enemy);
            }
        }
    }
}
