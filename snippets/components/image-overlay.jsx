
export const ImageOverlay = ({
    baseImage,
    overlayImage,
    baseAlt = "Base image",
    overlayAlt = "Overlay image",
    overlayPosition = "center",
    overlayWidth = "80%",
    overlayHeight = "100%",
    overlayTop,
    overlayLeft,
    overlayRight,
    overlayBottom,
    offsetX = "0px",
    offsetY = "0px",
    borderRadius = "8px",
    className = ""
}) => {
    const getPositionStyles = () => {
        const styles = {};

        // If custom position values are provided, use them
        if (overlayTop !== undefined) styles.top = overlayTop;
        if (overlayLeft !== undefined) styles.left = overlayLeft;
        if (overlayRight !== undefined) styles.right = overlayRight;
        if (overlayBottom !== undefined) styles.bottom = overlayBottom;

        // Otherwise use preset positions
        if (!overlayTop && !overlayLeft && !overlayRight && !overlayBottom) {
            switch (overlayPosition) {
                case "center":
                    styles.top = `calc(50% + ${offsetY})`;
                    styles.left = `calc(50% + ${offsetX})`;
                    styles.transform = 'translate(-50%, -50%)';
                    break;
                case "top-left":
                    styles.top = offsetY;
                    styles.left = offsetX;
                    break;
                case "top-right":
                    styles.top = offsetY;
                    styles.right = offsetX;
                    break;
                case "bottom-left":
                    styles.bottom = offsetY;
                    styles.left = offsetX;
                    break;
                case "bottom-right":
                    styles.bottom = offsetY;
                    styles.right = offsetX;
                    break;
                case "top-center":
                    styles.top = offsetY;
                    styles.left = `calc(50% + ${offsetX})`;
                    styles.transform = 'translateX(-50%)';
                    break;
                case "bottom-center":
                    styles.bottom = offsetY;
                    styles.left = `calc(50% + ${offsetX})`;
                    styles.transform = 'translateX(-50%)';
                    break;
                default:
                    styles.top = `calc(50% + ${offsetY})`;
                    styles.left = `calc(50% + ${offsetX})`;
                    styles.transform = 'translate(-50%, -50%)';
            }
        } else if (offsetX !== "0px" || offsetY !== "0px") {
            // Apply offsets even with custom positioning
            styles.transform = `translate(${offsetX}, ${offsetY})`;
        }

        return styles;
    };

    return (
        <div
            className={`relative overflow-hidden inline-block ${className}`}
            style={{
                borderRadius,
                transform: 'translateZ(0)', // Fixes compositing layer issues
                verticalAlign: 'top', // Fixes inline-block baseline gap
                lineHeight: 0 // Removes inline spacing
            }}
        >
            <img
                src={baseImage}
                alt={baseAlt}
                className="block w-full h-auto"
                style={{ borderRadius }}
            />
            <div
                className="absolute overflow-hidden"
                style={{
                    top: 0,
                    left: 0,
                    right: 0,
                    bottom: 0,
                    clipPath: 'inset(0)' // Additional clipping insurance
                }}
            >
                <div
                    className="absolute overflow-hidden"
                    style={{
                        ...getPositionStyles(),
                        width: overlayWidth,
                        height: overlayHeight
                    }}
                >
                    <img
                        src={overlayImage}
                        alt={overlayAlt}
                        className="block object-contain"
                        style={{
                            borderRadius,
                            maxWidth: '100%',
                            maxHeight: '100%',
                            width: 'auto',
                            height: 'auto'
                        }}
                    />
                </div>
            </div>
        </div>
    );
};
