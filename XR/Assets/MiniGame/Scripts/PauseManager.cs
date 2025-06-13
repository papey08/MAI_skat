using UnityEngine;

public class PauseManager : MonoBehaviour
{
    public GameObject pauseTextUI;
    public GameObject blackHole;

    private bool isPaused = false;
    private BlackHoleController blackHoleController;

    private void Start()
    {
        if (blackHole != null)
            blackHoleController = blackHole.GetComponent<BlackHoleController>();

        if (pauseTextUI != null)
            pauseTextUI.SetActive(false);
    }

    private void Update()
    {
        if (Input.GetKeyDown(KeyCode.P))
        {
            TogglePause();
        }
    }

    public void TogglePause()
    {
        isPaused = !isPaused;

        Time.timeScale = isPaused ? 0f : 1f;

        if (pauseTextUI != null)
            pauseTextUI.SetActive(isPaused);

        if (blackHoleController != null)
            blackHoleController.enabled = !isPaused;
    }
}
