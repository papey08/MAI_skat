using UnityEngine;

public class ControlPoint : MonoBehaviour
{
    public float activationDistance;
    public GameObject controlPointIndicator;
    public GameObject selectionMenu;

    private Transform player;

    void Start()
    {
        player = GameObject.FindGameObjectWithTag("Player")?.transform;

        if (controlPointIndicator != null)
            controlPointIndicator.SetActive(false);

        if (selectionMenu != null)
            selectionMenu.SetActive(false);

    }

    void Update()
    {
        if (player == null) return;

        float distance = Vector3.Distance(player.position, transform.position);
        if (distance <= activationDistance)
        {
            if (controlPointIndicator != null)
                controlPointIndicator.SetActive(true);

            if (Input.GetKeyDown(KeyCode.C) && selectionMenu != null)
            {
                selectionMenu.SetActive(true);
                Time.timeScale = 0f;
            }
        }
        else
        {
            if (controlPointIndicator != null)
                controlPointIndicator.SetActive(false);
        }
    }
}
