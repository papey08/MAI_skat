using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class RestartButtonManager : MonoBehaviour
{
    public Button restartButton;
    public GameObject missionFailedText;
    public GameObject playerShip;

    public EnemySpawner[] spawners;

    void Start()
    {
        restartButton.gameObject.SetActive(false);
        missionFailedText.SetActive(false);
        restartButton.onClick.AddListener(RestartGame);
    }

    void Update()
    {
        if (playerShip == null)
        {
            ShowMissionFailedText();
            ShowRestartButton();
        }
    }

    public void ShowRestartButton()
    {
        restartButton.gameObject.SetActive(true);
    }

    public void ShowMissionFailedText()
    {
        missionFailedText.SetActive(true);
    }

    void RestartGame()
    {
        SceneManager.LoadScene(SceneManager.GetActiveScene().name);
    }
}
