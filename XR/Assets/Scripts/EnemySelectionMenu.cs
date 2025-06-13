using UnityEngine;
using UnityEngine.UI;

public class EnemySelectionMenu : MonoBehaviour
{
    public Image enemyIcon;
    public Sprite[] enemySprites;
    public GameObject[] spawners;
    private int currentIndex = 0;

    public Button leftButton;
    public Button rightButton;
    public Button confirmButton;

    void Start()
    {
        UpdateUI();
        leftButton.onClick.AddListener(PreviousEnemy);
        rightButton.onClick.AddListener(NextEnemy);
        confirmButton.onClick.AddListener(ConfirmSelection);
    }

    void UpdateUI()
    {
        enemyIcon.sprite = enemySprites[currentIndex];
    }

    void PreviousEnemy()
    {
        currentIndex = (currentIndex - 1 + enemySprites.Length) % enemySprites.Length;
        UpdateUI();
    }

    void NextEnemy()
    {
        currentIndex = (currentIndex + 1) % enemySprites.Length;
        UpdateUI();
    }

    void ConfirmSelection()
    {
        for (int i = 0; i < spawners.Length; i++)
        {
            EnemySpawner spawner = spawners[i].GetComponent<EnemySpawner>();
            if (spawner != null)
                spawner.isActive = i == currentIndex;
        }

        gameObject.SetActive(false);
        Time.timeScale = 1f;
    }
}
