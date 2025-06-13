using UnityEngine;

public class AbsorbableObject : MonoBehaviour
{
    public float Scale => transform.localScale.x;

    public GameObject levelCompleteUI;

    private bool levelCompleted = false;

    private void Update()
    {
        CheckAbsorption();
    }

    private void CheckAbsorption()
    {
        if (levelCompleted) return;

        GameObject blackHole = GameObject.FindWithTag("Player");
        if (blackHole == null) return;

        BlackHoleController bh = blackHole.GetComponent<BlackHoleController>();
        float bhRadius = bh.GetVisualRadius();

        float distance = Vector2.Distance(transform.position, blackHole.transform.position);
        float thisRadius = GetComponent<SpriteRenderer>().bounds.size.x / 2f;

        if (distance <= bhRadius + thisRadius && blackHole.transform.localScale.x > Scale)
        {
            if (CompareTag("LevelGoal"))
            {
                levelCompleted = true;

                if (levelCompleteUI != null)
                {
                    levelCompleteUI.SetActive(true);
                }
            }

            Absorb(bh);
        }
    }

    private void Absorb(BlackHoleController bh)
    {
        bh.Absorb(GetComponent<SpriteRenderer>().bounds.size.x / 2f);
        Destroy(gameObject);
    }
}
