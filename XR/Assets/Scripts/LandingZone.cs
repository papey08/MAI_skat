using UnityEngine;
using UnityEngine.UI;
using UnityEngine.SceneManagement;

public class LandingZone : MonoBehaviour
{
    public float detectionRadius;
    public float requiredHoldTime;
    public RawImage landingIndicator;
    public GameObject returnToMenuButton;
    public GameObject missionSuccessText;

    private float holdTimer = 0f;
    private bool isPaused = false;
    private Transform player;

    private Color lightGreen = new Color(0.6f, 1f, 0.6f, 1f);
    private Color darkGreen = new Color(0f, 0.5f, 0f, 1f);

    void Start()
    {
        player = GameObject.FindGameObjectWithTag("Player")?.transform;

        if (landingIndicator != null)
            landingIndicator.gameObject.SetActive(false);

        if (returnToMenuButton != null)
            returnToMenuButton.SetActive(false);

        if (missionSuccessText != null)
            missionSuccessText.SetActive(false);
    }

    void Update()
    {
        if (player == null || isPaused) return;

        float distance = Vector3.Distance(transform.position, player.position);
        bool isInZone = distance <= detectionRadius;

        if (isInZone)
        {
            if (landingIndicator != null && !landingIndicator.gameObject.activeSelf)
                landingIndicator.gameObject.SetActive(true);

            if (Input.GetKey(KeyCode.L))
            {
                holdTimer += Time.deltaTime;
                SetIndicatorColor(darkGreen);

                if (holdTimer >= requiredHoldTime)
                    PauseGame();
            }
            else
            {
                holdTimer = 0f;
                SetIndicatorColor(lightGreen);
            }
        }
        else
        {
            holdTimer = 0f;

            if (landingIndicator != null && landingIndicator.gameObject.activeSelf)
                landingIndicator.gameObject.SetActive(false);
        }
    }

    void SetIndicatorColor(Color color)
    {
        if (landingIndicator != null)
            landingIndicator.color = color;
    }

    void PauseGame()
    {
        isPaused = true;
        Time.timeScale = 0f;

        if (landingIndicator != null)
            landingIndicator.gameObject.SetActive(false);

        if (returnToMenuButton != null)
        {
            returnToMenuButton.SetActive(true);
            Button btn = returnToMenuButton.GetComponent<Button>();
            if (btn != null)
            {
                btn.onClick.RemoveAllListeners();
                btn.onClick.AddListener(ReturnToMainMenu);
            }
        }

        if (missionSuccessText != null)
            missionSuccessText.SetActive(true);
    }

    void ReturnToMainMenu()
    {
        Time.timeScale = 1f;
        SceneManager.LoadScene("Scenes/MainMenu");
    }
}
