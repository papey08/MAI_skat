using UnityEngine;
using UnityEngine.SceneManagement;

public class MiniGameTrigger : MonoBehaviour
{
    public string miniGameSceneName;
    public float triggerRadius;

    public Transform player;

    private bool bossIsDead = false;

    void OnEnable()
    {
        BossDeathNotifier.OnBossDeath += OnBossDefeated;
    }

    void OnDisable()
    {
        BossDeathNotifier.OnBossDeath -= OnBossDefeated;
    }

    void Update()
    {
        if (!bossIsDead || player == null) return;

        float distance = Vector3.Distance(transform.position, player.position);
        if (distance <= triggerRadius)
        {
            SceneManager.LoadScene(miniGameSceneName);
        }
    }

    void OnBossDefeated()
    {
        bossIsDead = true;
    }
}
